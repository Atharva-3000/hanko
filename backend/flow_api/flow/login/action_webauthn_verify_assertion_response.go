package login

import (
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	auditlog "github.com/teamhanko/hanko/backend/audit_log"
	"github.com/teamhanko/hanko/backend/flow_api/flow/shared"
	"github.com/teamhanko/hanko/backend/flow_api/services"
	"github.com/teamhanko/hanko/backend/flowpilot"
	"github.com/teamhanko/hanko/backend/persistence/models"
)

type WebauthnVerifyAssertionResponse struct {
	shared.Action
}

func (a WebauthnVerifyAssertionResponse) GetName() flowpilot.ActionName {
	return shared.ActionWebauthnVerifyAssertionResponse
}

func (a WebauthnVerifyAssertionResponse) GetDescription() string {
	return "Send the result which was generated by using a webauthn credential."
}

func (a WebauthnVerifyAssertionResponse) Initialize(c flowpilot.InitializationContext) {
	if !c.Stash().Get("webauthn_available").Bool() {
		c.SuspendAction()
	}

	// We have to check for 'preflight' (hardcoded so as to not introduce dependency on capabilities package)
	// because at the time of the response/schema generation for the 'login_init' state the flow has not actually
	// progressed to that state yet (i.e. it is still in the 'preflight' state).
	if c.CurrentStateEquals("preflight") {
		if !c.Stash().Get("webauthn_conditional_mediation_available").Bool() {
			c.SuspendAction()
		}
	}

	c.AddInputs(flowpilot.JSONInput("assertion_response").Required(true).Persist(false))
}

func (a WebauthnVerifyAssertionResponse) Execute(c flowpilot.ExecutionContext) error {
	deps := a.GetDeps(c)

	if valid := c.ValidateInputData(); !valid {
		return c.ContinueFlowWithError(c.GetCurrentState(), flowpilot.ErrorFormDataInvalid)
	}

	if !c.Stash().Get("webauthn_session_data_id").Exists() {
		return errors.New("webauthn_session_data_id is not present in the stash")
	}

	sessionDataID := uuid.FromStringOrNil(c.Stash().Get("webauthn_session_data_id").String())
	assertionResponse := c.Input().Get("assertion_response").String()

	params := services.VerifyAssertionResponseParams{
		Tx:                deps.Tx,
		SessionDataID:     sessionDataID,
		AssertionResponse: assertionResponse,
	}

	userModel, err := deps.WebauthnService.VerifyAssertionResponse(params)
	if err != nil {
		if errors.Is(err, services.ErrInvalidWebauthnCredential) {
			err = deps.AuditLogger.CreateWithConnection(
				deps.Tx,
				deps.HttpContext,
				models.AuditLogLoginFailure,
				userModel,
				err,
				auditlog.Detail("login_method", "passkey"),
				auditlog.Detail("flow_id", c.GetFlowID()))

			if err != nil {
				return fmt.Errorf("could not create audit log: %w", err)
			}

			return c.ContinueFlowWithError(shared.StateLoginInit, shared.ErrorPasskeyInvalid.Wrap(err))
		}

		return fmt.Errorf("failed to verify assertion response: %w", err)
	}

	err = c.Stash().Set("user_id", userModel.ID.String())
	if err != nil {
		return fmt.Errorf("failed to set user_id to the stash: %w", err)
	}

	// Set only for audit logging purposes.
	err = c.Stash().Set("login_method", "passkey")
	if err != nil {
		return fmt.Errorf("failed to set login_method to the stash: %w", err)
	}

	return c.ContinueFlow(shared.StateSuccess)
}

func (a WebauthnVerifyAssertionResponse) Finalize(c flowpilot.FinalizationContext) error {
	return nil
}
