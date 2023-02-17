package helper

import (
	"net/http"

	"github.com/naeemaei/golang-clean-web-api/pkg/service_errors"
)

var StatusCodeMapping = map[string]int{

	// OTP
	service_errors.OptExists:   409,
	service_errors.OtpUsed:     409,
	service_errors.OtpNotValid: 400,

	// User
	service_errors.EmailExists:      409,
	service_errors.UsernameExists:   409,
	service_errors.RecordNotFound:   404,
	service_errors.PermissionDenied: 403,
}

func TranslateErrorToStatusCode(err error) int {
	value, ok := StatusCodeMapping[err.Error()]
	if !ok {
		return http.StatusInternalServerError
	}
	return value
}
