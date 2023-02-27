package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/naeemaei/golang-clean-web-api/api/dto"
	_ "github.com/naeemaei/golang-clean-web-api/api/helper"
	"github.com/naeemaei/golang-clean-web-api/config"
	"github.com/naeemaei/golang-clean-web-api/services"
)

type CompanyHandler struct {
	service *services.CompanyService
}

func NewCompanyHandler(cfg *config.Config) *CompanyHandler {
	return &CompanyHandler{service: services.NewCompanyService(cfg)}
}

// CreateCompany godoc
// @Summary Create a Company
// @Description Create a Company
// @Tags Companies
// @Accept json
// @produces json
// @Param Request body dto.CreateCompanyRequest true "Create a Company"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CompanyResponse}
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/companies/ [post]
// @Security AuthBearer
func (h *CompanyHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

// UpdateCompany godoc
// @Summary Update a Company
// @Description Update a Company
// @Tags Companies
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCompanyRequest true "Update a Company"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CompanyResponse}
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/companies/{id} [put]
// @Security AuthBearer
func (h *CompanyHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

// DeleteCompany godoc
// @Summary Delete a Company
// @Description Delete a Company
// @Tags Companies
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 201 {object} helper.BaseHttpResponse "response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/companies/{id} [delete]
// @Security AuthBearer
func (h *CompanyHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)
}

// GetCompany godoc
// @Summary Get a Company
// @Description Get a Company
// @Tags Companies
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CompanyResponse}
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/companies/{id} [get]
// @Security AuthBearer
func (h *CompanyHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

// GetCompany godoc
// @Summary get a Company
// @Description get a Company
// @Tags Companies
// @Accept  json
// @Produce  json
// @Param Request body dto.PaginationInputWithFilter true "request body"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.CompanyResponse],validationErrors=any}
// @Failure 400 {object} helper.BaseHttpResponse "Description"
// @Failure 401 {object} helper.BaseHttpResponse{result=any,validationErrors=any} "Description"
// @Router /v1/companies/get-by-filter [post]
// @Security AuthBearer
func (h *CompanyHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFilter)
}
