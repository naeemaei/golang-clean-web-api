package helper

type ResultCode int

const (
	ValidationError ResultCode = 4001
	OtpLimiterError ResultCode = 4002
	CustomRecovery  ResultCode = 5000
	InternalError   ResultCode = 5001
	LimiterError    ResultCode = 5002
	AuthError       ResultCode = 5003
	Success         ResultCode = 0
)