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

type CarModelColorHandler struct {
	service *services.CarModelColorService
}

func NewCarModelColorHandler(cfg *config.Config) *CarModelColorHandler {
	return &CarModelColorHandler{service: services.NewCarModelColorService(cfg)}
}

// CreateCarModelColor godoc
// @Summary Create a CarModelColor
// @Description Create a CarModelColor
// @Tags CarModelColors
// @Accept json
// @produces json
// @Param Request body dto.CreateUpdateCarModelColorRequest true "Create a CarModelColor"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CarModelColorResponse}
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-model-colors/ [post]
// @Security AuthBearer
func (h *CarModelColorHandler) Create(c *gin.Context) {
	req := dto.CreateUpdateCarModelColorRequest{}
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

// UpdateCarModelColor godoc
// @Summary Update a CarModelColor
// @Description Update a CarModelColor
// @Tags CarModelColors
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.CreateUpdateCarModelColorRequest true "Update a CarModelColor"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CarModelColorResponse}
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-model-colors/{id} [put]
// @Security AuthBearer
func (h *CarModelColorHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	req := dto.CreateUpdateCarModelColorRequest{}
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

// DeleteCarModelColor godoc
// @Summary Delete a CarModelColor
// @Description Delete a CarModelColor
// @Tags CarModelColors
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 201 {object} helper.BaseHttpResponse "response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-model-colors/{id} [delete]
// @Security AuthBearer
func (h *CarModelColorHandler) Delete(c *gin.Context) {
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

// GetCarModelColor godoc
// @Summary Get a CarModelColor
// @Description Get a CarModelColor
// @Tags CarModelColors
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarModelColorResponse}
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-model-colors/{id} [get]
// @Security AuthBearer
func (h *CarModelColorHandler) GetById(c *gin.Context) {
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

// GetCarModelColor godoc
// @Summary get a CarModelColor
// @Description get a CarModelColor
// @Tags CarModelColors
// @Accept  json
// @Produce  json
// @Param Request body dto.PaginationInputWithFilter true "request body"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.CarModelColorResponse],validationErrors=any}
// @Failure 400 {object} helper.BaseHttpResponse "Description"
// @Failure 401 {object} helper.BaseHttpResponse{result=any,validationErrors=any} "Description"
// @Router /v1/car-model-colors/get-by-filter [post]
// @Security AuthBearer
func (h *CarModelColorHandler) GetByFilter(c *gin.Context) {

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
