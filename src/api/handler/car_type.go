package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/naeemaei/golang-clean-web-api/api/dto"
	_ "github.com/naeemaei/golang-clean-web-api/api/helper"
	"github.com/naeemaei/golang-clean-web-api/config"
	"github.com/naeemaei/golang-clean-web-api/dependency"
	_ "github.com/naeemaei/golang-clean-web-api/domain/filter"
	"github.com/naeemaei/golang-clean-web-api/usecase"
)

type CarTypeHandler struct {
	usecase *usecase.CarTypeUsecase
}

func NewCarTypeHandler(cfg *config.Config) *CarTypeHandler {
	return &CarTypeHandler{
		usecase: usecase.NewCarTypeUsecase(cfg, dependency.GetCarTypeRepository(cfg)),
	}
}

// CreateCarType godoc
// @Summary Create a CarType
// @Description Create a CarType
// @Tags CarTypes
// @Accept json
// @produces json
// @Param Request body dto.CreateCarTypeRequest true "Create a CarType"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CarTypeResponse} "CarType response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-types/ [post]
// @Security AuthBearer
func (h *CarTypeHandler) Create(c *gin.Context) {
	Create(c, dto.ToCreateCarType, dto.ToCarTypeResponse, h.usecase.Create)
}

// UpdateCarType godoc
// @Summary Update a CarType
// @Description Update a CarType
// @Tags CarTypes
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCarTypeRequest true "Update a CarType"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarTypeResponse} "CarType response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/car-types/{id} [put]
// @Security AuthBearer
func (h *CarTypeHandler) Update(c *gin.Context) {
	Update(c, dto.ToUpdateCarType, dto.ToCarTypeResponse, h.usecase.Update)
}

// DeleteCarType godoc
// @Summary Delete a CarType
// @Description Delete a CarType
// @Tags CarTypes
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse "response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/car-types/{id} [delete]
// @Security AuthBearer
func (h *CarTypeHandler) Delete(c *gin.Context) {
	Delete(c, h.usecase.Delete)
}

// GetCarType godoc
// @Summary Get a CarType
// @Description Get a CarType
// @Tags CarTypes
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarTypeResponse} "CarType response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/car-types/{id} [get]
// @Security AuthBearer
func (h *CarTypeHandler) GetById(c *gin.Context) {
	GetById(c, dto.ToCarTypeResponse, h.usecase.GetById)
}

// GetCarTypes godoc
// @Summary Get CarTypes
// @Description Get CarTypes
// @Tags CarTypes
// @Accept json
// @produces json
// @Param Request body filter.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=filter.PagedList[dto.CarTypeResponse]} "CarType response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-types/get-by-filter [post]
// @Security AuthBearer
func (h *CarTypeHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, dto.ToCarTypeResponse, h.usecase.GetByFilter)
}
