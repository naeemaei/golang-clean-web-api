package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/naeemaei/golang-clean-web-api/api/middlewares"
	"github.com/naeemaei/golang-clean-web-api/api/routers"
	validation "github.com/naeemaei/golang-clean-web-api/api/validations"
	"github.com/naeemaei/golang-clean-web-api/config"
	"github.com/naeemaei/golang-clean-web-api/docs"
	"github.com/naeemaei/golang-clean-web-api/pkg/logging"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitServer(cfg *config.Config) {
	r := gin.New()

	RegisterValidators()

	r.Use(middlewares.DefaultStructuredLogger(cfg))
	r.Use(middlewares.Cors(cfg))
	r.Use(gin.Logger(), gin.Recovery() /*middlewares.TestMiddleware()*/, middlewares.LimitByRequest())

	RegisterRoutes(r)
	RegisterSwagger(r, cfg)
	logger := logging.NewLogger(cfg)
	logger.Info(logging.General, logging.Startup, "Started", nil)
	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}

func RegisterRoutes(r *gin.Engine) {
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
}

func RegisterValidators() {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("mobile", validation.IranianMobileNumberValidator, true)
		val.RegisterValidation("password", validation.PasswordValidator, true)
	}
}

func RegisterSwagger(r *gin.Engine, cfg *config.Config) {
	docs.SwaggerInfo.Title = "golang web api"
	docs.SwaggerInfo.Description = "golang web api"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", cfg.Server.Port)
	docs.SwaggerInfo.Schemes = []string{"http"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
