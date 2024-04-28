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

type CarModelImageHandler struct {
	usecase *usecase.CarModelImageUsecase
}

func NewCarModelImageHandler(cfg *config.Config) *CarModelImageHandler {
	return &CarModelImageHandler{
		usecase: usecase.NewCarModelImageUsecase(cfg, dependency.GetCarModelImageRepository(cfg)),
	}
}

// CreateCarModelImage godoc
// @Summary Create a CarModelImage
// @Description Create a CarModelImage
// @Tags CarModelImages
// @Accept json
// @produces json
// @Param Request body dto.CreateCarModelImageRequest true "Create a CarModelImage"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CarModelImageResponse} "CarModelImage response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-model-images/ [post]
// @Security AuthBearer
func (h *CarModelImageHandler) Create(c *gin.Context) {
	Create(c, dto.ToCreateCarModelImage, dto.ToCarModelImageResponse, h.usecase.Create)
}

// UpdateCarModelImage godoc
// @Summary Update a CarModelImage
// @Description Update a CarModelImage
// @Tags CarModelImages
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCarModelImageRequest true "Update a CarModelImage"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarModelImageResponse} "CarModelImage response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/car-model-images/{id} [put]
// @Security AuthBearer
func (h *CarModelImageHandler) Update(c *gin.Context) {
	Update(c, dto.ToUpdateCarModelImage, dto.ToCarModelImageResponse, h.usecase.Update)
}

// DeleteCarModelImage godoc
// @Summary Delete a CarModelImage
// @Description Delete a CarModelImage
// @Tags CarModelImages
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse "response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/car-model-images/{id} [delete]
// @Security AuthBearer
func (h *CarModelImageHandler) Delete(c *gin.Context) {
	Delete(c, h.usecase.Delete)
}

// GetCarModelImage godoc
// @Summary Get a CarModelImage
// @Description Get a CarModelImage
// @Tags CarModelImages
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarModelImageResponse} "CarModelImage response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/car-model-images/{id} [get]
// @Security AuthBearer
func (h *CarModelImageHandler) GetById(c *gin.Context) {
	GetById(c, dto.ToCarModelImageResponse, h.usecase.GetById)
}

// GetCarModelImages godoc
// @Summary Get CarModelImages
// @Description Get CarModelImages
// @Tags CarModelImages
// @Accept json
// @produces json
// @Param Request body filter.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=filter.PagedList[dto.CarModelImageResponse]} "CarModelImage response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-model-images/get-by-filter [post]
// @Security AuthBearer
func (h *CarModelImageHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, dto.ToCarModelImageResponse, h.usecase.GetByFilter)
}
