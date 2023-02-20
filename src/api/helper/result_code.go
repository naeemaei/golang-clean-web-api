package helper

type ResultCode int

const (
	ValidationError ResultCode = 40001
	LimiterError    ResultCode = 42901
	OtpLimiterError ResultCode = 42902
	CustomRecovery  ResultCode = 50001
	InternalError   ResultCode = 50002
	AuthError       ResultCode = 40101
	Success         ResultCode = 0
)
