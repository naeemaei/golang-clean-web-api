package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/naeemaei/golang-clean-web-api/api/dto"
	_ "github.com/naeemaei/golang-clean-web-api/api/helper"
	"github.com/naeemaei/golang-clean-web-api/config"
	"github.com/naeemaei/golang-clean-web-api/dependency"
	_ "github.com/naeemaei/golang-clean-web-api/domain/filter"
	"github.com/naeemaei/golang-clean-web-api/usecase"
)

type PersianYearHandler struct {
	usecase *usecase.PersianYearUsecase
}

func NewPersianYearHandler(cfg *config.Config) *PersianYearHandler {
	return &PersianYearHandler{
		usecase: usecase.NewPersianYearUsecase(cfg, dependency.GetPersianYearRepository(cfg)),
	}
}

// CreatePersianYear godoc
// @Summary Create a PersianYear
// @Description Create a PersianYear
// @Tags PersianYears
// @Accept json
// @produces json
// @Param Request body dto.CreatePersianYearRequest true "Create a PersianYear"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.PersianYearResponse} "PersianYear response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/years/ [post]
// @Security AuthBearer
func (h *PersianYearHandler) Create(c *gin.Context) {

	Create(c, dto.ToCreatePersianYear, dto.ToPersianYearResponse, h.usecase.Create)
}

// UpdatePersianYear godoc
// @Summary Update a PersianYear
// @Description Update a PersianYear
// @Tags PersianYears
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdatePersianYearRequest true "Update a PersianYear"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PersianYearResponse} "PersianYear response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/years/{id} [put]
// @Security AuthBearer
func (h *PersianYearHandler) Update(c *gin.Context) {
	Update(c, dto.ToUpdatePersianYear, dto.ToPersianYearResponse, h.usecase.Update)
}

// DeletePersianYear godoc
// @Summary Delete a PersianYear
// @Description Delete a PersianYear
// @Tags PersianYears
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse "response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/years/{id} [delete]
// @Security AuthBearer
func (h *PersianYearHandler) Delete(c *gin.Context) {
	Delete(c, h.usecase.Delete)
}

// GetPersianYear godoc
// @Summary Get a PersianYear
// @Description Get a PersianYear
// @Tags PersianYears
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PersianYearResponse} "PersianYear response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/years/{id} [get]
// @Security AuthBearer
func (h *PersianYearHandler) GetById(c *gin.Context) {
	GetById(c, dto.ToPersianYearResponse, h.usecase.GetById)
}

// GetPersianYears godoc
// @Summary Get PersianYears
// @Description Get PersianYears
// @Tags PersianYears
// @Accept json
// @produces json
// @Param Request body filter.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=filter.PagedList[dto.PersianYearResponse]} "PersianYear response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/years/get-by-filter [post]
// @Security AuthBearer
func (h *PersianYearHandler) GetByFilter(c *gin.Context) {

	GetByFilter(c, dto.ToPersianYearResponse, h.usecase.GetByFilter)
}
