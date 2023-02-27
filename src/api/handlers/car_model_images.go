package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/naeemaei/golang-clean-web-api/api/dto"
	_ "github.com/naeemaei/golang-clean-web-api/api/helper"
	"github.com/naeemaei/golang-clean-web-api/config"
	"github.com/naeemaei/golang-clean-web-api/services"
)

type CarModelImageHandler struct {
	service *services.CarModelImageService
}

func NewCarModelImageHandler(cfg *config.Config) *CarModelImageHandler {
	return &CarModelImageHandler{service: services.NewCarModelImageService(cfg)}
}

// CreateCarModelImage godoc
// @Summary Create a CarModelImage
// @Description Create a CarModelImage
// @Tags CarModelImages
// @Accept json
// @produces json
// @Param Request body dto.CreateUpdateCarModelImageRequest true "Create a CarModelImage"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CarModelImageResponse}
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-model-images/ [post]
// @Security AuthBearer
func (h *CarModelImageHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

// UpdateCarModelImage godoc
// @Summary Update a CarModelImage
// @Description Update a CarModelImage
// @Tags CarModelImages
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.CreateUpdateCarModelImageRequest true "Update a CarModelImage"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CarModelImageResponse}
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-model-images/{id} [put]
// @Security AuthBearer
func (h *CarModelImageHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

// DeleteCarModelImage godoc
// @Summary Delete a CarModelImage
// @Description Delete a CarModelImage
// @Tags CarModelImages
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 201 {object} helper.BaseHttpResponse "response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-model-images/{id} [delete]
// @Security AuthBearer
func (h *CarModelImageHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)
}

// GetCarModelImage godoc
// @Summary Get a CarModelImage
// @Description Get a CarModelImage
// @Tags CarModelImages
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarModelImageResponse}
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-model-images/{id} [get]
// @Security AuthBearer
func (h *CarModelImageHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

// GetCarModelImage godoc
// @Summary get a CarModelImage
// @Description get a CarModelImage
// @Tags CarModelImages
// @Accept  json
// @Produce  json
// @Param Request body dto.PaginationInputWithFilter true "request body"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.CarModelImageResponse],validationErrors=any}
// @Failure 400 {object} helper.BaseHttpResponse "Description"
// @Failure 401 {object} helper.BaseHttpResponse{result=any,validationErrors=any} "Description"
// @Router /v1/car-model-images/get-by-filter [post]
// @Security AuthBearer
func (h *CarModelImageHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFilter)
}
