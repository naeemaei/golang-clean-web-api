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

type PropertyHandler struct {
	service *services.PropertyService
}

func NewPropertyHandler(cfg *config.Config) *PropertyHandler {
	return &PropertyHandler{service: services.NewPropertyService(cfg)}
}

// CreateProperty godoc
// @Summary Create a Property
// @Description Create a Property
// @Tags Properties
// @Accept json
// @produces json
// @Param Request body dto.CreateUpdatePropertyRequest true "Create a Property"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.PropertyResponse}
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/properties/ [post]
// @Security AuthBearer
func (h *PropertyHandler) Create(c *gin.Context) {
	req := dto.CreateUpdatePropertyRequest{}
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

// UpdateProperty godoc
// @Summary Update a Property
// @Description Update a Property
// @Tags Properties
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.CreateUpdatePropertyRequest true "Update a Property"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.PropertyResponse}
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/properties/{id} [put]
// @Security AuthBearer
func (h *PropertyHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	req := dto.CreateUpdatePropertyRequest{}
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

// DeleteProperty godoc
// @Summary Delete a Property
// @Description Delete a Property
// @Tags Properties
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 201 {object} helper.BaseHttpResponse "response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/properties/{id} [delete]
// @Security AuthBearer
func (h *PropertyHandler) Delete(c *gin.Context) {
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

// GetProperty godoc
// @Summary Get a Property
// @Description Get a Property
// @Tags Properties
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PropertyResponse}
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/properties/{id} [get]
// @Security AuthBearer
func (h *PropertyHandler) GetById(c *gin.Context) {
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

// GetProperty godoc
// @Summary get a Property
// @Description get a Property
// @Tags Properties
// @Accept  json
// @Produce  json
// @Param Request body dto.PaginationInputWithFilter true "request body"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.PropertyResponse],validationErrors=any}
// @Failure 400 {object} helper.BaseHttpResponse "Description"
// @Failure 401 {object} helper.BaseHttpResponse{result=any,validationErrors=any} "Description"
// @Router /v1/properties/get-by-filter [post]
// @Security AuthBearer
func (h *PropertyHandler) GetByFilter(c *gin.Context) {

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
