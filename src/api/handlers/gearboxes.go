package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/naeemaei/golang-clean-web-api/api/dto"
	_ "github.com/naeemaei/golang-clean-web-api/api/helper"
	"github.com/naeemaei/golang-clean-web-api/config"
	"github.com/naeemaei/golang-clean-web-api/services"
)

type GearboxHandler struct {
	service *services.GearboxService
}

func NewGearboxHandler(cfg *config.Config) *GearboxHandler {
	return &GearboxHandler{service: services.NewGearboxService(cfg)}
}

// CreateGearbox godoc
// @Summary Create a country
// @Description Create a country
// @Tags Gearboxes
// @Accept json
// @produces json
// @Param Request body dto.CreateGearboxRequest true "Create a country"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.GearboxResponse}
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/gearboxes/ [post]
// @Security AuthBearer
func (h *GearboxHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

// UpdateGearbox godoc
// @Summary Update a country
// @Description Update a country
// @Tags Gearboxes
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateGearboxRequest true "Update a country"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.GearboxResponse} "Gearbox response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/gearboxes/{id} [put]
// @Security AuthBearer
func (h *GearboxHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

// DeleteGearbox godoc
// @Summary Delete a country
// @Description Delete a country
// @Tags Gearboxes
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse "response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/gearboxes/{id} [delete]
// @Security AuthBearer
func (h *GearboxHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)
}

// GetGearbox godoc
// @Summary Get a country
// @Description Get a country
// @Tags Gearboxes
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.GearboxResponse}
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/gearboxes/{id} [get]
// @Security AuthBearer
func (h *GearboxHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

// GetGearboxes godoc
// @Summary Get Gearboxes
// @Description Get Gearboxes
// @Tags Gearboxes
// @Accept json
// @produces json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.GearboxResponse]} "Gearbox response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/gearboxes/get-by-filter [post]
// @Security AuthBearer
func (h *GearboxHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFilter)
}
