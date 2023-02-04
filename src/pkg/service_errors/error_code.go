package service_errors

const (
	// Token
	UnExpectedError = "Expected error"
	ClaimsNotFound  = "Claims not found"

	// OTP
	OptExists   = "Otp exists"
	OtpUsed     = "Otp used"
	OtpNotValid = "Otp invalid"

	// User
	EmailExists    = "Email exists"
	UsernameExists = "Username exists"
)
