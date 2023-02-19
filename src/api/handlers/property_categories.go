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

type PropertyCategoryHandler struct {
	service *services.PropertyCategoryService
}

func NewPropertyCategoryHandler(cfg *config.Config) *PropertyCategoryHandler {
	return &PropertyCategoryHandler{service: services.NewPropertyCategoryService(cfg)}
}

// CreatePropertyCategory godoc
// @Summary Create a PropertyCategory
// @Description Create a PropertyCategory
// @Tags PropertyCategories
// @Accept json
// @produces json
// @Param Request body dto.CreateUpdatePropertyCategoryRequest true "Create a PropertyCategory"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.PropertyCategoryResponse}
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/property-categories/ [post]
// @Security AuthBearer
func (h *PropertyCategoryHandler) Create(c *gin.Context) {
	req := dto.CreateUpdatePropertyCategoryRequest{}
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

// UpdatePropertyCategory godoc
// @Summary Update a PropertyCategory
// @Description Update a PropertyCategory
// @Tags PropertyCategories
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.CreateUpdatePropertyCategoryRequest true "Update a PropertyCategory"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.PropertyCategoryResponse}
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/property-categories/{id} [put]
// @Security AuthBearer
func (h *PropertyCategoryHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	req := dto.CreateUpdatePropertyCategoryRequest{}
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

// DeletePropertyCategory godoc
// @Summary Delete a PropertyCategory
// @Description Delete a PropertyCategory
// @Tags PropertyCategories
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 201 {object} helper.BaseHttpResponse "response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/property-categories/{id} [delete]
// @Security AuthBearer
func (h *PropertyCategoryHandler) Delete(c *gin.Context) {
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

// GetPropertyCategory godoc
// @Summary Get a PropertyCategory
// @Description Get a PropertyCategory
// @Tags PropertyCategories
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PropertyCategoryResponse}
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/property-categories/{id} [get]
// @Security AuthBearer
func (h *PropertyCategoryHandler) GetById(c *gin.Context) {
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

// GetPropertyCategory godoc
// @Summary get a PropertyCategory
// @Description get a PropertyCategory
// @Tags PropertyCategories
// @Accept  json
// @Produce  json
// @Param Request body dto.PaginationInputWithFilter true "request body"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.PropertyCategoryResponse],validationErrors=any}
// @Failure 400 {object} helper.BaseHttpResponse "Description"
// @Failure 401 {object} helper.BaseHttpResponse{result=any,validationErrors=any} "Description"
// @Router /v1/property-categories/get-by-filter [post]
// @Security AuthBearer
func (h *PropertyCategoryHandler) GetByFilter(c *gin.Context) {

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
