package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/naeemaei/golang-clean-web-api/api/helper"
	"github.com/naeemaei/golang-clean-web-api/config"
	constant "github.com/naeemaei/golang-clean-web-api/constant"
	"github.com/naeemaei/golang-clean-web-api/pkg/service_errors"
	"github.com/naeemaei/golang-clean-web-api/usecase"
)

func Authentication(cfg *config.Config) gin.HandlerFunc {
	var tokenUsecase = usecase.NewTokenUsecase(cfg)

	return func(c *gin.Context) {
		var err error
		claimMap := map[string]interface{}{}
		auth := c.GetHeader(constant.AuthorizationHeaderKey)
		token := strings.Split(auth, " ")
		if auth == "" || len(token) < 2 {
			err = &service_errors.ServiceError{EndUserMessage: service_errors.TokenRequired}
		} else {
			claimMap, err = tokenUsecase.GetClaims(token[1])
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					err = &service_errors.ServiceError{EndUserMessage: service_errors.TokenExpired}
				default:
					err = &service_errors.ServiceError{EndUserMessage: service_errors.TokenInvalid}
				}
			}
		}
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helper.GenerateBaseResponseWithError(
				nil, false, helper.AuthError, err,
			))
			return
		}

		c.Set(constant.UserIdKey, claimMap[constant.UserIdKey])
		c.Set(constant.FirstNameKey, claimMap[constant.FirstNameKey])
		c.Set(constant.LastNameKey, claimMap[constant.LastNameKey])
		c.Set(constant.UsernameKey, claimMap[constant.UsernameKey])
		c.Set(constant.EmailKey, claimMap[constant.EmailKey])
		c.Set(constant.MobileNumberKey, claimMap[constant.MobileNumberKey])
		c.Set(constant.RolesKey, claimMap[constant.RolesKey])
		c.Set(constant.ExpireTimeKey, claimMap[constant.ExpireTimeKey])

		c.Next()
	}
}

func Authorization(validRoles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.Keys) == 0 {
			c.AbortWithStatusJSON(http.StatusForbidden, helper.GenerateBaseResponse(nil, false, helper.ForbiddenError))
			return
		}
		rolesVal := c.Keys[constant.RolesKey]
		fmt.Println(rolesVal)
		if rolesVal == nil {
			c.AbortWithStatusJSON(http.StatusForbidden, helper.GenerateBaseResponse(nil, false, helper.ForbiddenError))
			return
		}
		roles := rolesVal.([]interface{})
		val := map[string]int{}
		for _, item := range roles {
			val[item.(string)] = 0
		}

		for _, item := range validRoles {
			if _, ok := val[item]; ok {
				c.Next()
				return
			}
		}
		c.AbortWithStatusJSON(http.StatusForbidden, helper.GenerateBaseResponse(nil, false, helper.ForbiddenError))
	}
}
