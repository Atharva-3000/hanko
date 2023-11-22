package passcode

import (
	"errors"
	"github.com/gofrs/uuid"
	"github.com/teamhanko/hanko/backend/flow_api/shared"
	"github.com/teamhanko/hanko/backend/flowpilot"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var maxPasscodeTries = 3

type VerifyPasscode struct {
	shared.Action
}

func (a VerifyPasscode) GetName() flowpilot.ActionName {
	return ActionVerifyPasscode
}

func (a VerifyPasscode) GetDescription() string {
	return "Enter a passcode."
}

func (a VerifyPasscode) Initialize(c flowpilot.InitializationContext) {
	c.AddInputs(flowpilot.StringInput("code").Required(true))
}

func (a VerifyPasscode) Execute(c flowpilot.ExecutionContext) error {
	deps := a.GetDeps(c)

	if valid := c.ValidateInputData(); !valid {
		return c.ContinueFlowWithError(c.GetCurrentState(), flowpilot.ErrorFormDataInvalid)
	}

	passcodeId, err := uuid.FromString(c.Stash().Get("passcode_id").String())
	if err != nil {
		return err
	}

	passcode, err := deps.Persister.GetPasscodePersister().Get(passcodeId)
	if err != nil {
		return err
	}
	if passcode == nil {
		return errors.New("passcode not found")
	}

	expirationTime := passcode.CreatedAt.Add(time.Duration(passcode.Ttl) * time.Second)
	if expirationTime.Before(time.Now().UTC()) {
		return c.ContinueFlowWithError(c.GetCurrentState(), flowpilot.ErrorFormDataInvalid.Wrap(errors.New("passcode is expired")))
	}

	err = bcrypt.CompareHashAndPassword([]byte(passcode.Code), []byte(c.Input().Get("code").String()))
	if err != nil {
		passcode.TryCount += 1
		if passcode.TryCount >= maxPasscodeTries {
			err = deps.Persister.GetPasscodePersister().Delete(*passcode)
			if err != nil {
				return err
			}
			err = c.Stash().Delete("passcode_id")
			if err != nil {
				return err
			}

			return c.ContinueFlowWithError(c.GetCurrentState(), shared.ErrorPasscodeMaxAttemptsReached)
		}
		return c.ContinueFlowWithError(c.GetCurrentState(), shared.ErrorPasscodeInvalid.Wrap(err))
	}

	// !?
	//err = c.Stash().Set("user_id", passcode.UserId)
	//if err != nil {
	//	return fmt.Errorf("failed to set user_id to the stash: %w", err)
	//}

	if !c.Stash().Get("user_id").Exists() {
		return c.ContinueFlowWithError(c.GetErrorState(), flowpilot.ErrorOperationNotPermitted.Wrap(errors.New("account does not exist")))
	}

	err = c.Stash().Set("email_verified", true) // TODO: maybe change attribute path
	if err != nil {
		return err
	}

	err = deps.Persister.GetPasscodePersisterWithConnection(deps.Tx).Delete(*passcode)
	if err != nil {
		return err
	}

	return c.EndSubFlow()
}
