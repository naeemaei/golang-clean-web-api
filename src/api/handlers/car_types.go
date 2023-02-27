package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/naeemaei/golang-clean-web-api/api/dto"
	_ "github.com/naeemaei/golang-clean-web-api/api/helper"
	"github.com/naeemaei/golang-clean-web-api/config"
	"github.com/naeemaei/golang-clean-web-api/services"
)

type CarTypeHandler struct {
	service *services.CarTypeService
}

func NewCarTypeHandler(cfg *config.Config) *CarTypeHandler {
	return &CarTypeHandler{service: services.NewCarTypeService(cfg)}
}

// CreateCarType godoc
// @Summary Create a CarType
// @Description Create a CarType
// @Tags CarTypes
// @Accept json
// @produces json
// @Param Request body dto.CreateCarTypeRequest true "Create a CarType"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CarTypeResponse}
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-types/ [post]
// @Security AuthBearer
func (h *CarTypeHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

// UpdateCarType godoc
// @Summary Update a CarType
// @Description Update a CarType
// @Tags CarTypes
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCarTypeRequest true "Update a CarType"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CarTypeResponse}
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-types/{id} [put]
// @Security AuthBearer
func (h *CarTypeHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

// DeleteCarType godoc
// @Summary Delete a CarType
// @Description Delete a CarType
// @Tags CarTypes
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 201 {object} helper.BaseHttpResponse "response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-types/{id} [delete]
// @Security AuthBearer
func (h *CarTypeHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)
}

// GetCarType godoc
// @Summary Get a CarType
// @Description Get a CarType
// @Tags CarTypes
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarTypeResponse}
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-types/{id} [get]
// @Security AuthBearer
func (h *CarTypeHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

// GetCarType godoc
// @Summary get a CarType
// @Description get a CarType
// @Tags CarTypes
// @Accept  json
// @Produce  json
// @Param Request body dto.PaginationInputWithFilter true "request body"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.CarTypeResponse],validationErrors=any}
// @Failure 400 {object} helper.BaseHttpResponse "Description"
// @Failure 401 {object} helper.BaseHttpResponse{result=any,validationErrors=any} "Description"
// @Router /v1/car-types/get-by-filter [post]
// @Security AuthBearer
func (h *CarTypeHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFilter)
}
