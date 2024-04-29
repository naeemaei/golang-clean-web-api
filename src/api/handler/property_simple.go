package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/naeemaei/golang-clean-web-api/api/dto"
	"github.com/naeemaei/golang-clean-web-api/api/helper"
	"github.com/naeemaei/golang-clean-web-api/config"
	"github.com/naeemaei/golang-clean-web-api/dependency"
	"github.com/naeemaei/golang-clean-web-api/domain/filter"
	"github.com/naeemaei/golang-clean-web-api/usecase"
)

type PropertySimpleHandler struct {
	usecase *usecase.PropertyUsecase
}

func NewPropertySimpleHandler(cfg *config.Config) *PropertySimpleHandler {
	return &PropertySimpleHandler{
		usecase: usecase.NewPropertyUsecase(cfg, dependency.GetPropertyRepository(cfg)),
	}
}

// CreateProperty godoc
// @Summary Create a Property
// @Description Create a Property
// @Tags Properties
// @Accept json
// @produces json
// @Param Request body dto.CreatePropertyRequest true "Create a Property"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.PropertyResponse} "Property response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/properties/ [post]
// @Security AuthBearer
func (h *PropertySimpleHandler) Create(c *gin.Context) {
	request := new(dto.CreatePropertyRequest)
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}

	// map http request body to usecase input and call use case method
	property, err := h.usecase.Create(c, dto.ToCreateProperty(*request))

	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	// map usecase response to http response
	response := dto.ToPropertyResponse(property)

	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(response, true, 0))
}

// UpdateProperty godoc
// @Summary Update a Property
// @Description Update a Property
// @Tags Properties
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdatePropertyRequest true "Update a Property"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PropertyResponse} "Property response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/properties/{id} [put]
// @Security AuthBearer
func (h *PropertySimpleHandler) Update(c *gin.Context) {
	// bind http request
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	request := new(dto.UpdatePropertyRequest)
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return

	}
	// map http request body to usecase input and call use case method
	property, err := h.usecase.Update(c, id, dto.ToUpdateProperty(*request))

	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	// map usecase response to http response
	response := dto.ToPropertyResponse(property)

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(response, true, 0))
}

// DeleteProperty godoc
// @Summary Delete a Property
// @Description Delete a Property
// @Tags Properties
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse "response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/properties/{id} [delete]
// @Security AuthBearer
func (h *PropertySimpleHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	if id == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound,
			helper.GenerateBaseResponse(nil, false, helper.ValidationError))
		return
	}

	err := h.usecase.Delete(c, id)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
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
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PropertyResponse} "Property response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/properties/{id} [get]
// @Security AuthBearer
func (h *PropertySimpleHandler) GetById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	if id == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound,
			helper.GenerateBaseResponse(nil, false, helper.ValidationError))
		return
	}

	// call use case method
	property, err := h.usecase.GetById(c, id)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	// map usecase response to http response
	response := dto.ToPropertyResponse(property)

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(response, true, 0))
}

// GetProperties godoc
// @Summary Get Properties
// @Description Get Properties
// @Tags Properties
// @Accept json
// @produces json
// @Param Request body filter.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=filter.PagedList[dto.PropertyResponse]} "Property response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/properties/get-by-filter [post]
// @Security AuthBearer
func (h *PropertySimpleHandler) GetByFilter(c *gin.Context) {
	req := new(filter.PaginationInputWithFilter)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}

	// call use case method
	properties, err := h.usecase.GetByFilter(c, *req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}
	response := filter.PagedList[dto.PropertyResponse]{
		PageNumber:      properties.PageNumber,
		PageSize:        properties.PageSize,
		TotalRows:       properties.TotalRows,
		TotalPages:      properties.TotalPages,
		HasPreviousPage: properties.HasPreviousPage,
		HasNextPage:     properties.HasNextPage,
	}

	// map usecase response to http response
	items := []dto.PropertyResponse{}
	for _, item := range *properties.Items {

		items = append(items, dto.ToPropertyResponse(item))
	}
	response.Items = &items

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(response, true, 0))
}
