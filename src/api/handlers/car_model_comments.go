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

type CarModelCommentHandler struct {
	service *services.CarModelCommentService
}

func NewCarModelCommentHandler(cfg *config.Config) *CarModelCommentHandler {
	return &CarModelCommentHandler{service: services.NewCarModelCommentService(cfg)}
}

// CreateCarModelComment godoc
// @Summary Create a CarModelComment
// @Description Create a CarModelComment
// @Tags CarModelComments
// @Accept json
// @produces json
// @Param Request body dto.CreateUpdateCarModelCommentRequest true "Create a CarModelComment"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CarModelCommentResponse}
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-model-comments/ [post]
// @Security AuthBearer
func (h *CarModelCommentHandler) Create(c *gin.Context) {
	req := dto.CreateUpdateCarModelCommentRequest{}
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

// UpdateCarModelComment godoc
// @Summary Update a CarModelComment
// @Description Update a CarModelComment
// @Tags CarModelComments
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.CreateUpdateCarModelCommentRequest true "Update a CarModelComment"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CarModelCommentResponse}
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-model-comments/{id} [put]
// @Security AuthBearer
func (h *CarModelCommentHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	req := dto.CreateUpdateCarModelCommentRequest{}
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

// DeleteCarModelComment godoc
// @Summary Delete a CarModelComment
// @Description Delete a CarModelComment
// @Tags CarModelComments
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 201 {object} helper.BaseHttpResponse "response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-model-comments/{id} [delete]
// @Security AuthBearer
func (h *CarModelCommentHandler) Delete(c *gin.Context) {
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

// GetCarModelComment godoc
// @Summary Get a CarModelComment
// @Description Get a CarModelComment
// @Tags CarModelComments
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarModelCommentResponse}
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-model-comments/{id} [get]
// @Security AuthBearer
func (h *CarModelCommentHandler) GetById(c *gin.Context) {
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

// GetCarModelComment godoc
// @Summary get a CarModelComment
// @Description get a CarModelComment
// @Tags CarModelComments
// @Accept  json
// @Produce  json
// @Param Request body dto.PaginationInputWithFilter true "request body"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.CarModelCommentResponse],validationErrors=any}
// @Failure 400 {object} helper.BaseHttpResponse "Description"
// @Failure 401 {object} helper.BaseHttpResponse{result=any,validationErrors=any} "Description"
// @Router /v1/car-model-comments/get-by-filter [post]
// @Security AuthBearer
func (h *CarModelCommentHandler) GetByFilter(c *gin.Context) {

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
