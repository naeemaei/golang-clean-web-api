package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/naeemaei/golang-clean-web-api/api/middleware"
	"github.com/naeemaei/golang-clean-web-api/api/router"
	validation "github.com/naeemaei/golang-clean-web-api/api/validation"
	"github.com/naeemaei/golang-clean-web-api/config"
	"github.com/naeemaei/golang-clean-web-api/docs"
	"github.com/naeemaei/golang-clean-web-api/pkg/logging"
	"github.com/naeemaei/golang-clean-web-api/pkg/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var logger = logging.NewLogger(config.GetConfig())

func InitServer(cfg *config.Config) {
	gin.SetMode(cfg.Server.RunMode)
	r := gin.New()
	RegisterValidators()
	RegisterPrometheus()

	r.Use(middleware.DefaultStructuredLogger(cfg))
	r.Use(middleware.Cors(cfg))
	r.Use(middleware.Prometheus())
	r.Use(gin.Logger(), gin.CustomRecovery(middleware.ErrorHandler) /*middleware.TestMiddleware()*/, middleware.LimitByRequest())

	RegisterRoutes(r, cfg)
	RegisterSwagger(r, cfg)
	logger := logging.NewLogger(cfg)
	logger.Info(logging.General, logging.Startup, "Started", nil)
	err := r.Run(fmt.Sprintf(":%s", cfg.Server.InternalPort))
	if err != nil {
		logger.Fatal(logging.General, logging.Startup, err.Error(), nil)
	}
}

func RegisterRoutes(r *gin.Engine, cfg *config.Config) {
	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		// Test
		health := v1.Group("/health")
		testRouter := v1.Group("/test" /*middleware.Authentication(cfg), middleware.Authorization([]string{"admin"})*/)

		// User
		users := v1.Group("/users")

		// Base
		countries := v1.Group("/countries", middleware.Authentication(cfg), middleware.Authorization([]string{"admin"}))
		cities := v1.Group("/cities", middleware.Authentication(cfg), middleware.Authorization([]string{"admin"}))
		files := v1.Group("/files", middleware.Authentication(cfg), middleware.Authorization([]string{"admin"}))
		companies := v1.Group("/companies", middleware.Authentication(cfg), middleware.Authorization([]string{"admin"}))
		colors := v1.Group("/colors", middleware.Authentication(cfg), middleware.Authorization([]string{"admin"}))
		years := v1.Group("/years", middleware.Authentication(cfg), middleware.Authorization([]string{"admin"}))

		// Property
		properties := v1.Group("/properties", middleware.Authentication(cfg), middleware.Authorization([]string{"admin"}))
		propertyCategories := v1.Group("/property-categories", middleware.Authentication(cfg), middleware.Authorization([]string{"admin"}))

		// Car
		carTypes := v1.Group("/car-types", middleware.Authentication(cfg), middleware.Authorization([]string{"admin"}))
		gearboxes := v1.Group("/gearboxes", middleware.Authentication(cfg), middleware.Authorization([]string{"admin"}))
		carModels := v1.Group("/car-models", middleware.Authentication(cfg), middleware.Authorization([]string{"admin"}))
		carModelColors := v1.Group("/car-model-colors", middleware.Authentication(cfg), middleware.Authorization([]string{"admin"}))
		carModelYears := v1.Group("/car-model-years", middleware.Authentication(cfg), middleware.Authorization([]string{"admin"}))
		carModelPriceHistories := v1.Group("/car-model-price-histories", middleware.Authentication(cfg), middleware.Authorization([]string{"admin"}))
		carModelImages := v1.Group("/car-model-images", middleware.Authentication(cfg), middleware.Authorization([]string{"admin"}))
		carModelProperties := v1.Group("/car-model-properties", middleware.Authentication(cfg), middleware.Authorization([]string{"admin"}))
		carModelComments := v1.Group("/car-model-comments", middleware.Authentication(cfg), middleware.Authorization([]string{"admin", "default"}))

		// Test
		router.Health(health)
		router.TestRouter(testRouter)

		// User
		router.User(users, cfg)

		// Base
		router.Country(countries, cfg)
		router.City(cities, cfg)
		router.File(files, cfg)
		router.Company(companies, cfg)
		router.Color(colors, cfg)
		router.Year(years, cfg)

		// Property
		router.Property(properties, cfg)
		router.PropertyCategory(propertyCategories, cfg)

		// Car
		router.CarType(carTypes, cfg)
		router.Gearbox(gearboxes, cfg)
		router.CarModel(carModels, cfg)
		router.CarModelColor(carModelColors, cfg)
		router.CarModelYear(carModelYears, cfg)
		router.CarModelPriceHistory(carModelPriceHistories, cfg)
		router.CarModelImage(carModelImages, cfg)
		router.CarModelProperty(carModelProperties, cfg)
		router.CarModelComment(carModelComments, cfg)

		r.Static("/static", "./uploads")

		r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	}

	v2 := api.Group("/v2")
	{
		health := v2.Group("/health")
		router.Health(health)
	}
}

func RegisterValidators() {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		err := val.RegisterValidation("mobile", validation.IranianMobileNumberValidator, true)
		if err != nil {
			logger.Error(logging.Validation, logging.Startup, err.Error(), nil)
		}
		err = val.RegisterValidation("password", validation.PasswordValidator, true)
		if err != nil {
			logger.Error(logging.Validation, logging.Startup, err.Error(), nil)
		}
	}
}

func RegisterSwagger(r *gin.Engine, cfg *config.Config) {
	docs.SwaggerInfo.Title = "golang web api"
	docs.SwaggerInfo.Description = "golang web api"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", cfg.Server.ExternalPort)
	docs.SwaggerInfo.Schemes = []string{"http"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func RegisterPrometheus() {
	err := prometheus.Register(metrics.DbCall)
	if err != nil {
		logger.Error(logging.Prometheus, logging.Startup, err.Error(), nil)
	}

	err = prometheus.Register(metrics.HttpDuration)
	if err != nil {
		logger.Error(logging.Prometheus, logging.Startup, err.Error(), nil)
	}
}
