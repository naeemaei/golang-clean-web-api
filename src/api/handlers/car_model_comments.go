package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/naeemaei/golang-clean-web-api/api/dto"
	_ "github.com/naeemaei/golang-clean-web-api/api/helper"
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
// @Param Request body dto.CreateCarModelCommentRequest true "Create a CarModelComment"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CarModelCommentResponse}
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-model-comments/ [post]
// @Security AuthBearer
func (h *CarModelCommentHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

// UpdateCarModelComment godoc
// @Summary Update a CarModelComment
// @Description Update a CarModelComment
// @Tags CarModelComments
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCarModelCommentRequest true "Update a CarModelComment"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CarModelCommentResponse}
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-model-comments/{id} [put]
// @Security AuthBearer
func (h *CarModelCommentHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
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
	Delete(c, h.service.Delete)
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
	GetById(c, h.service.GetById)
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
	GetByFilter(c, h.service.GetByFilter)
}
