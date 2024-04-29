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

type CarModelColorHandler struct {
	usecase *usecase.CarModelColorUsecase
}

func NewCarModelColorHandler(cfg *config.Config) *CarModelColorHandler {
	return &CarModelColorHandler{
		usecase: usecase.NewCarModelColorUsecase(cfg, dependency.GetCarModelColorRepository(cfg)),
	}
}

// CreateCarModelColor godoc
// @Summary Create a CarModelColor
// @Description Create a CarModelColor
// @Tags CarModelColors
// @Accept json
// @produces json
// @Param Request body dto.CreateCarModelColorRequest true "Create a CarModelColor"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CarModelColorResponse} "CarModelColor response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-model-colors/ [post]
// @Security AuthBearer
func (h *CarModelColorHandler) Create(c *gin.Context) {
	Create(c, dto.ToCreateCarModelColor, dto.ToCarModelColorResponse, h.usecase.Create)
}

// UpdateCarModelColor godoc
// @Summary Update a CarModelColor
// @Description Update a CarModelColor
// @Tags CarModelColors
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCarModelColorRequest true "Update a CarModelColor"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarModelColorResponse} "CarModelColor response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/car-model-colors/{id} [put]
// @Security AuthBearer
func (h *CarModelColorHandler) Update(c *gin.Context) {
	Update(c, dto.ToUpdateCarModelColor, dto.ToCarModelColorResponse, h.usecase.Update)
}

// DeleteCarModelColor godoc
// @Summary Delete a CarModelColor
// @Description Delete a CarModelColor
// @Tags CarModelColors
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse "response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/car-model-colors/{id} [delete]
// @Security AuthBearer
func (h *CarModelColorHandler) Delete(c *gin.Context) {
	Delete(c, h.usecase.Delete)
}

// GetCarModelColor godoc
// @Summary Get a CarModelColor
// @Description Get a CarModelColor
// @Tags CarModelColors
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarModelColorResponse} "CarModelColor response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/car-model-colors/{id} [get]
// @Security AuthBearer
func (h *CarModelColorHandler) GetById(c *gin.Context) {
	GetById(c, dto.ToCarModelColorResponse, h.usecase.GetById)
}

// GetCarModelColors godoc
// @Summary Get CarModelColors
// @Description Get CarModelColors
// @Tags CarModelColors
// @Accept json
// @produces json
// @Param Request body filter.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=filter.PagedList[dto.CarModelColorResponse]} "CarModelColor response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-model-colors/get-by-filter [post]
// @Security AuthBearer
func (h *CarModelColorHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, dto.ToCarModelColorResponse, h.usecase.GetByFilter)
}
