package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/naeemaei/golang-clean-web-api/api/routers"
	"github.com/naeemaei/golang-clean-web-api/config"
)

func InitServer() {
	cfg := config.GetConfig()
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	
	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		health := v1.Group("/health")
		test_router := v1.Group("/test")

		routers.Health(health)
		routers.TestRouter(test_router)
	}

	v2 := api.Group("/v2")
	{
		health := v2.Group("/health")
		routers.Health(health)
	}

	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}
