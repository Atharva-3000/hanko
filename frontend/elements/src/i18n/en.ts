import { Translation } from "./translations";

export const en: Translation = {
  headlines: {
    error: "An error has occurred",
    loginEmail: "Sign in or create account",
    loginEmailNoSignup: "Sign in",
    loginFinished: "Login successful",
    loginPasscode: "Enter passcode",
    loginPassword: "Enter password",
    registerAuthenticator: "Create a passkey",
    registerConfirm: "Create account?",
    registerPassword: "Set new password",
    profileEmails: "Emails",
    profilePassword: "Password",
    profilePasskeys: "Passkeys",
    isPrimaryEmail: "Primary email address",
    setPrimaryEmail: "Set primary email address",
    createEmail: "Enter a new email",
    createUsername: "Enter a new username",
    emailVerified: "Verified",
    emailUnverified: "Unverified",
    emailDelete: "Delete",
    renamePasskey: "Rename passkey",
    deletePasskey: "Delete passkey",
    lastUsedAt: "Last used at",
    createdAt: "Created at",
    connectedAccounts: "Connected accounts",
    deleteAccount: "Delete account",
    accountNotFound: "Account not found",
    signIn: "Sign in",
    signUp: "Create account",
    selectLoginMethod: "Select login method",
    setupLoginMethod: "Set up login method",
    lastUsed: "Last seen",
    ipAddress: "IP address",
    revokeSession: "Revoke session",
    profileSessions: "Sessions"
  },
  texts: {
    enterPasscode: 'Enter the passcode that was sent to "{emailAddress}".',
    enterPasscodeNoEmail:
      "Enter the passcode that was sent to your primary email address.",
    setupPasskey:
      "Sign in to your account easily and securely with a passkey. Note: Your biometric data is only stored on your devices and will never be shared with anyone.",
    createAccount:
      'No account exists for "{emailAddress}". Do you want to create a new account?',
    passwordFormatHint:
      "Must be between {minLength} and {maxLength} characters long.",
    setPrimaryEmail: "Set this email address to be used for contacting you.",
    isPrimaryEmail:
      "This email address will be used to contact you if necessary.",
    emailVerified: "This email address has been verified.",
    emailUnverified: "This email address has not been verified.",
    emailDelete:
      "If you delete this email address, it can no longer be used to sign in.",
    renamePasskey: "Set a name for the passkey.",
    deletePasskey: "Delete this passkey from your account.",
    deleteAccount:
      "Are you sure you want to delete this account? All data will be deleted immediately and cannot be recovered.",
    noAccountExists: 'No account exists for "{emailAddress}".',
    selectLoginMethodForFutureLogins:
      "Select one of the following login methods to use for future logins.",
    howDoYouWantToLogin: "How do you want to login?",
  },
  labels: {
    or: "or",
    no: "no",
    yes: "yes",
    email: "Email",
    continue: "Continue",
    skip: "Skip",
    save: "Save",
    password: "Password",
    passkey: "Passkey",
    passcode: "Passcode",
    signInPassword: "Sign in with a password",
    signInPasscode: "Sign in with a passcode",
    forgotYourPassword: "Forgot your password?",
    back: "Back",
    signInPasskey: "Sign in with a passkey",
    registerAuthenticator: "Create a passkey",
    signIn: "Sign in",
    signUp: "Create account",
    sendNewPasscode: "Send new code",
    passwordRetryAfter: "Retry in {passwordRetryAfter}",
    passcodeResendAfter: "Request a new code in {passcodeResendAfter}",
    unverifiedEmail: "unverified",
    primaryEmail: "primary",
    setAsPrimaryEmail: "Set as primary",
    verify: "Verify",
    delete: "Delete",
    newEmailAddress: "New email address",
    newPassword: "New password",
    rename: "Rename",
    newPasskeyName: "New passkey name",
    addEmail: "Add email",
    createPasskey: "Create a passkey",
    webauthnUnsupported: "Passkeys are not supported by your browser",
    signInWith: "Sign in with {provider}",
    deleteAccount: "Yes, delete this account.",
    emailOrUsername: "Email or username",
    username: "Username",
    optional: "optional",
    dontHaveAnAccount: "Don't have an account?",
    alreadyHaveAnAccount: "Already have an account?",
    changeUsername: "Change username",
    setUsername: "Set username",
    changePassword: "Change password",
    setPassword: "Set password",
    revoke: "Revoke",
    currentSession: "Current session",
  },
  errors: {
    somethingWentWrong:
      "A technical error has occurred. Please try again later.",
    requestTimeout: "The request timed out.",
    invalidPassword: "Wrong email or password.",
    invalidPasscode: "The passcode provided was not correct.",
    passcodeAttemptsReached:
      "The passcode was entered incorrectly too many times. Please request a new code.",
    tooManyRequests:
      "Too many requests have been made. Please wait to repeat the requested operation.",
    unauthorized: "Your session has expired. Please log in again.",
    invalidWebauthnCredential: "This passkey cannot be used anymore.",
    passcodeExpired: "The passcode has expired. Please request a new one.",
    userVerification:
      "User verification required. Please ensure your authenticator device is protected with a PIN or biometric.",
    emailAddressAlreadyExistsError: "The email address already exists.",
    maxNumOfEmailAddressesReached: "No further email addresses can be added.",
    thirdPartyAccessDenied:
      "Access denied. The request was cancelled by the user or the provider has denied access for other reasons.",
    thirdPartyMultipleAccounts:
      "Cannot identify account. The email address is used by multiple accounts.",
    thirdPartyUnverifiedEmail:
      "Email verification required. Please verify the used email address with your provider.",
    signupDisabled: "Account registration is disabled.",
  },
  flowErrors: {
    technical_error: "A technical error has occurred. Please try again later.",
    flow_expired_error:
      "The session has expired, please click the button to restart.",
    value_invalid_error: "The entered value is invalid.",
    passcode_invalid: "The passcode provided was not correct.",
    passkey_invalid: "This passkey cannot be used anymore",
    passcode_max_attempts_reached:
      "The passcode was entered incorrectly too many times. Please request a new code.",
    rate_limit_exceeded:
      "Too many requests have been made. Please wait to repeat the requested operation.",
    unknown_username_error: "The username is unknown.",
    username_already_exists: "The username is already taken.",
    invalid_username_error:
      "The username must contain only letters, numbers, and underscores.",
    email_already_exists: "The email is already taken.",
    not_found: "The requested resource was not found.",
    operation_not_permitted_error: "The operation is not permitted.",
    flow_discontinuity_error:
      "The process cannot be continued due to user settings or the provider's configuration.",
    form_data_invalid_error: "The submitted form data contains errors.",
    unauthorized: "Your session has expired. Please log in again.",
    value_missing_error: "The value is missing.",
    value_too_long_error: "Value is too long.",
    value_too_short_error: "The value is too short.",
  },
};
