package service_errors

const (
	// Token
	UnExpectedError     = "Expected error"
	ClaimsNotFound      = "Claims not found"
	TokenRequired       = "token required"
	TokenExpired        = "token expired"
	TokenInvalid        = "token invalid"
	InvalidRefreshToken = "invalid refresh token"

	// OTP
	OptExists   = "Otp exists"
	OtpUsed     = "Otp used"
	OtpNotValid = "Otp invalid"

	// User
	EmailExists               = "Email exists"
	UsernameExists            = "Username exists"
	PermissionDenied          = "Permission denied"
	UsernameOrPasswordInvalid = "username or password invalid"
	InvalidRolesFormat        = "invalid roles format"

	// DB
	RecordNotFound = "record not found"
)
