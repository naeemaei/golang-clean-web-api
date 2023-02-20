package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/naeemaei/golang-clean-web-api/api/dto"
	"github.com/naeemaei/golang-clean-web-api/api/helper"
	"github.com/naeemaei/golang-clean-web-api/config"
	"github.com/naeemaei/golang-clean-web-api/services"
)

type CarModelHandler struct {
	service *services.CarModelService
}

func NewCarModelHandler(cfg *config.Config) *CarModelHandler {
	return &CarModelHandler{service: services.NewCarModelService(cfg)}
}

// CreateCarModel godoc
// @Summary Create a CarModel
// @Description Create a CarModel
// @Tags CarModels
// @Accept json
// @produces json
// @Param Request body dto.CreateUpdateCarModelRequest true "Create a CarModel"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CarModelResponse}
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-models/ [post]
// @Security AuthBearer
func (h *CarModelHandler) Create(c *gin.Context) {
	req := dto.CreateUpdateCarModelRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, 121, err))
		return
	}

	res, err := h.service.Create(c, &req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, 121, err))
		return
	}
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(res, true, 0))
}

// UpdateCarModel godoc
// @Summary Update a CarModel
// @Description Update a CarModel
// @Tags CarModels
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.CreateUpdateCarModelRequest true "Update a CarModel"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CarModelResponse}
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-models/{id} [put]
// @Security AuthBearer
func (h *CarModelHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	req := dto.CreateUpdateCarModelRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, 121, err))
		return
	}

	res, err := h.service.Update(c, id, &req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, 121, err))
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, 0))
}

// DeleteCarModel godoc
// @Summary Delete a CarModel
// @Description Delete a CarModel
// @Tags CarModels
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 201 {object} helper.BaseHttpResponse "response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-models/{id} [delete]
// @Security AuthBearer
func (h *CarModelHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	if id == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound,
			helper.GenerateBaseResponse(nil, false, 121))
		return
	}

	err := h.service.Delete(c, id)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, 121, err))
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(nil, true, 0))
}

// GetCarModel godoc
// @Summary Get a CarModel
// @Description Get a CarModel
// @Tags CarModels
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarModelResponse}
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-models/{id} [get]
// @Security AuthBearer
func (h *CarModelHandler) GetById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	if id == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound,
			helper.GenerateBaseResponse(nil, false, 121))
		return
	}

	res, err := h.service.GetById(c, id)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, 121, err))
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, 0))
}

// GetCarModel godoc
// @Summary get a CarModel
// @Description get a CarModel
// @Tags CarModels
// @Accept  json
// @Produce  json
// @Param Request body dto.PaginationInputWithFilter true "request body"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.CarModelResponse],validationErrors=any}
// @Failure 400 {object} helper.BaseHttpResponse "Description"
// @Failure 401 {object} helper.BaseHttpResponse{result=any,validationErrors=any} "Description"
// @Router /v1/car-models/get-by-filter [post]
// @Security AuthBearer
func (h *CarModelHandler) GetByFilter(c *gin.Context) {

	filter := new(dto.PaginationInputWithFilter)
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}

	response, err := h.service.GetByFilter(c, filter)
	if err != nil {
		c.JSON(helper.TranslateErrorToStatusCode(err), helper.GenerateBaseResponseWithError(nil, false, -1, err))
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(response, true, 0))
}
