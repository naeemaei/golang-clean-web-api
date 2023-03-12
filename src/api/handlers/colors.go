package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/naeemaei/golang-clean-web-api/api/dto"
	_ "github.com/naeemaei/golang-clean-web-api/api/helper"
	"github.com/naeemaei/golang-clean-web-api/config"
	"github.com/naeemaei/golang-clean-web-api/services"
)

type ColorHandler struct {
	service *services.ColorService
}

func NewColorHandler(cfg *config.Config) *ColorHandler {
	return &ColorHandler{
		service: services.NewColorService(cfg),
	}
}

// CreateColor godoc
// @Summary Create a Color
// @Description Create a Color
// @Tags Colors
// @Accept json
// @produces json
// @Param Request body dto.CreateColorRequest true "Create a Color"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.ColorResponse} "Color response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/colors/ [post]
// @Security AuthBearer
func (h *ColorHandler) Create(c *gin.Context) {
	Create(c,h.service.Create)
}

// UpdateColor godoc
// @Summary Update a Color
// @Description Update a Color
// @Tags Colors
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateColorRequest true "Update a Color"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.ColorResponse} "Color response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/colors/{id} [put]
// @Security AuthBearer
func (h *ColorHandler) Update(c *gin.Context) {
	Update(c,h.service.Update)
}

// DeleteColor godoc
// @Summary Delete a Color
// @Description Delete a Color
// @Tags Colors
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse "response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/colors/{id} [delete]
// @Security AuthBearer
func (h *ColorHandler) Delete(c *gin.Context) {
	Delete(c,h.service.Delete)
}

// GetColor godoc
// @Summary Get a Color
// @Description Get a Color
// @Tags Colors
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.ColorResponse} "Color response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/colors/{id} [get]
// @Security AuthBearer
func (h *ColorHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

// GetColors godoc
// @Summary Get Colors
// @Description Get Colors
// @Tags Colors
// @Accept json
// @produces json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.ColorResponse]} "Color response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/colors/get-by-filter [post]
// @Security AuthBearer
func (h *ColorHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFilter)
}
