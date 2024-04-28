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

type CarModelPriceHistoryHandler struct {
	usecase *usecase.CarModelPriceHistoryUsecase
}

func NewCarModelPriceHistoryHandler(cfg *config.Config) *CarModelPriceHistoryHandler {
	return &CarModelPriceHistoryHandler{
		usecase: usecase.NewCarModelPriceHistoryUsecase(cfg, dependency.GetCarModelPriceHistoryRepository(cfg)),
	}
}

// CreateCarModelPriceHistory godoc
// @Summary Create a CarModelPriceHistory
// @Description Create a CarModelPriceHistory
// @Tags CarModelPriceHistories
// @Accept json
// @produces json
// @Param Request body dto.CreateCarModelPriceHistoryRequest true "Create a CarModelPriceHistory"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CarModelPriceHistoryResponse} "CarModelPriceHistory response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-model-price-histories/ [post]
// @Security AuthBearer
func (h *CarModelPriceHistoryHandler) Create(c *gin.Context) {
	Create(c, dto.ToCreateCarModelPriceHistory, dto.ToCarModelPriceHistoryResponse, h.usecase.Create)
}

// UpdateCarModelPriceHistory godoc
// @Summary Update a CarModelPriceHistory
// @Description Update a CarModelPriceHistory
// @Tags CarModelPriceHistories
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCarModelPriceHistoryRequest true "Update a CarModelPriceHistory"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarModelPriceHistoryResponse} "CarModelPriceHistory response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/car-model-price-histories/{id} [put]
// @Security AuthBearer
func (h *CarModelPriceHistoryHandler) Update(c *gin.Context) {
	Update(c, dto.ToUpdateCarModelPriceHistory, dto.ToCarModelPriceHistoryResponse, h.usecase.Update)
}

// DeleteCarModelPriceHistory godoc
// @Summary Delete a CarModelPriceHistory
// @Description Delete a CarModelPriceHistory
// @Tags CarModelPriceHistories
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse "response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/car-model-price-histories/{id} [delete]
// @Security AuthBearer
func (h *CarModelPriceHistoryHandler) Delete(c *gin.Context) {
	Delete(c, h.usecase.Delete)
}

// GetCarModelPriceHistory godoc
// @Summary Get a CarModelPriceHistory
// @Description Get a CarModelPriceHistory
// @Tags CarModelPriceHistories
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarModelPriceHistoryResponse} "CarModelPriceHistory response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/car-model-price-histories/{id} [get]
// @Security AuthBearer
func (h *CarModelPriceHistoryHandler) GetById(c *gin.Context) {
	GetById(c, dto.ToCarModelPriceHistoryResponse, h.usecase.GetById)
}

// GetCarModelPriceHistories godoc
// @Summary Get CarModelPriceHistories
// @Description Get CarModelPriceHistories
// @Tags CarModelPriceHistories
// @Accept json
// @produces json
// @Param Request body filter.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=filter.PagedList[dto.CarModelPriceHistoryResponse]} "CarModelPriceHistory response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-model-price-histories/get-by-filter [post]
// @Security AuthBearer
func (h *CarModelPriceHistoryHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, dto.ToCarModelPriceHistoryResponse, h.usecase.GetByFilter)
}
