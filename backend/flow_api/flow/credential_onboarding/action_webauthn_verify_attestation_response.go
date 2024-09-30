package credential_onboarding

import (
	"github.com/teamhanko/hanko/backend/flow_api/flow/registration"
	"github.com/teamhanko/hanko/backend/flow_api/flow/shared"
	"github.com/teamhanko/hanko/backend/flowpilot"
)

type WebauthnVerifyAttestationResponse struct {
	shared.Action
}

func (a WebauthnVerifyAttestationResponse) GetName() flowpilot.ActionName {
	return shared.ActionWebauthnVerifyAttestationResponse
}

func (a WebauthnVerifyAttestationResponse) GetDescription() string {
	return "Send the result which was generated by creating a webauthn credential."
}

func (a WebauthnVerifyAttestationResponse) Initialize(c flowpilot.InitializationContext) {
	if !c.Stash().Get(shared.StashPathWebauthnAvailable).Bool() {
		c.SuspendAction()
	}

	c.AddInputs(flowpilot.JSONInput("public_key"))
}

func (a WebauthnVerifyAttestationResponse) Execute(c flowpilot.ExecutionContext) error {
	if valid := c.ValidateInputData(); !valid {
		return c.Error(flowpilot.ErrorFormDataInvalid)
	}

	if err := c.ExecuteHook(shared.VerifyAttestationResponse{}); err != nil {
		return err
	}

	c.PreventRevert()

	if err := c.ExecuteHook(registration.ScheduleMFACreationStates{}); err != nil {
		return err
	}

	return c.Continue()
}
