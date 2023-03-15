package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/naeemaei/golang-clean-web-api/api/dto"
	_ "github.com/naeemaei/golang-clean-web-api/api/helper"
	"github.com/naeemaei/golang-clean-web-api/config"
	"github.com/naeemaei/golang-clean-web-api/services"
)

type CarModelPropertyHandler struct {
	service *services.CarModelPropertyService
}

func NewCarModelPropertyHandler(cfg *config.Config) *CarModelPropertyHandler {
	return &CarModelPropertyHandler{
		service: services.NewCarModelPropertyService(cfg),
	}
}

// CreateCarModelProperty godoc
// @Summary Create a CarModelProperty
// @Description Create a CarModelProperty
// @Tags CarModelProperties
// @Accept json
// @produces json
// @Param Request body dto.CreateCarModelPropertyRequest true "Create a CarModelProperty"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CarModelPropertyResponse} "CarModelProperty response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-model-properties/ [post]
// @Security AuthBearer
func (h *CarModelPropertyHandler) Create(c *gin.Context) {
	Create(c,h.service.Create)
}

// UpdateCarModelProperty godoc
// @Summary Update a CarModelProperty
// @Description Update a CarModelProperty
// @Tags CarModelProperties
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCarModelPropertyRequest true "Update a CarModelProperty"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarModelPropertyResponse} "CarModelProperty response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/car-model-properties/{id} [put]
// @Security AuthBearer
func (h *CarModelPropertyHandler) Update(c *gin.Context) {
	Update(c,h.service.Update)
}

// DeleteCarModelProperty godoc
// @Summary Delete a CarModelProperty
// @Description Delete a CarModelProperty
// @Tags CarModelProperties
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse "response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/car-model-properties/{id} [delete]
// @Security AuthBearer
func (h *CarModelPropertyHandler) Delete(c *gin.Context) {
	Delete(c,h.service.Delete)
}

// GetCarModelProperty godoc
// @Summary Get a CarModelProperty
// @Description Get a CarModelProperty
// @Tags CarModelProperties
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarModelPropertyResponse} "CarModelProperty response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/car-model-properties/{id} [get]
// @Security AuthBearer
func (h *CarModelPropertyHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

// GetCarModelProperties godoc
// @Summary Get CarModelProperties
// @Description Get CarModelProperties
// @Tags CarModelProperties
// @Accept json
// @produces json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.CarModelPropertyResponse]} "CarModelProperty response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-model-properties/get-by-filter [post]
// @Security AuthBearer
func (h *CarModelPropertyHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFilter)
}
