package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/naeemaei/golang-clean-web-api/api/dto"
	_ "github.com/naeemaei/golang-clean-web-api/api/helper"
	"github.com/naeemaei/golang-clean-web-api/config"
	"github.com/naeemaei/golang-clean-web-api/services"
)

type CarModelPriceHistoryHandler struct {
	service *services.CarModelPriceHistoryService
}

func NewCarModelPriceHistoryHandler(cfg *config.Config) *CarModelPriceHistoryHandler {
	return &CarModelPriceHistoryHandler{service: services.NewCarModelPriceHistoryService(cfg)}
}

// CreateCarModelPriceHistory godoc
// @Summary Create a CarModelPriceHistory
// @Description Create a CarModelPriceHistory
// @Tags CarModelPriceHistories
// @Accept json
// @produces json
// @Param Request body dto.CreateCarModelPriceHistoryRequest true "Create a CarModelPriceHistory"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CarModelPriceHistoryResponse}
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-model-price-histories/ [post]
// @Security AuthBearer
func (h *CarModelPriceHistoryHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

// UpdateCarModelPriceHistory godoc
// @Summary Update a CarModelPriceHistory
// @Description Update a CarModelPriceHistory
// @Tags CarModelPriceHistories
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCarModelPriceHistoryRequest true "Update a CarModelPriceHistory"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CarModelPriceHistoryResponse}
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-model-price-histories/{id} [put]
// @Security AuthBearer
func (h *CarModelPriceHistoryHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

// DeleteCarModelPriceHistory godoc
// @Summary Delete a CarModelPriceHistory
// @Description Delete a CarModelPriceHistory
// @Tags CarModelPriceHistories
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 201 {object} helper.BaseHttpResponse "response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-model-price-histories/{id} [delete]
// @Security AuthBearer
func (h *CarModelPriceHistoryHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)
}

// GetCarModelPriceHistory godoc
// @Summary Get a CarModelPriceHistory
// @Description Get a CarModelPriceHistory
// @Tags CarModelPriceHistories
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarModelPriceHistoryResponse}
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-model-price-histories/{id} [get]
// @Security AuthBearer
func (h *CarModelPriceHistoryHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

// GetCarModelPriceHistory godoc
// @Summary get a CarModelPriceHistory
// @Description get a CarModelPriceHistory
// @Tags CarModelPriceHistories
// @Accept  json
// @Produce  json
// @Param Request body dto.PaginationInputWithFilter true "request body"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.CarModelPriceHistoryResponse],validationErrors=any}
// @Failure 400 {object} helper.BaseHttpResponse "Description"
// @Failure 401 {object} helper.BaseHttpResponse{result=any,validationErrors=any} "Description"
// @Router /v1/car-model-price-histories/get-by-filter [post]
// @Security AuthBearer
func (h *CarModelPriceHistoryHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFilter)
}
