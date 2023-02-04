package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naeemaei/golang-clean-web-api/api/helper"
)

func ErrorHandler(c *gin.Context, err any) {
	if err, ok := err.(error); ok {
		httpResponse := helper.GenerateBaseResponseWithError(nil, false, -500, err.(error))
		c.AbortWithStatusJSON(http.StatusInternalServerError, httpResponse)
		return
	}
	httpResponse := helper.GenerateBaseResponseWithAnyError(nil, false, -500, err)
	c.AbortWithStatusJSON(http.StatusInternalServerError, httpResponse)
}
