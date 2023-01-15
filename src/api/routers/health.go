package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/naeemaei/golang-clean-web-api/api/handlers"
)

func Health (r *gin.RouterGroup){
	handler := handlers.NewHealthHandler()

	r.GET("/", handler.Health) 
}