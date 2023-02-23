package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/naeemaei/golang-clean-web-api/api/handlers"
	"github.com/naeemaei/golang-clean-web-api/config"
)

func Country(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewCountryHandler(cfg)

	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
	r.POST("/get-by-filter", h.GetByFilter)
}
