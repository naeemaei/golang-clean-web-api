package helper

type ResultCode int

const (
	Success         ResultCode = 0
	ValidationError ResultCode = 40001
	AuthError       ResultCode = 40101
	ForbiddenError       ResultCode = 40301
	NotFoundError ResultCode = 40401
	LimiterError    ResultCode = 42901
	OtpLimiterError ResultCode = 42902
	CustomRecovery  ResultCode = 50001
	InternalError   ResultCode = 50002
)
