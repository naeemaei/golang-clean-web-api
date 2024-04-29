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

type CompanyHandler struct {
	usecase *usecase.CompanyUsecase
}

func NewCompanyHandler(cfg *config.Config) *CompanyHandler {
	return &CompanyHandler{
		usecase: usecase.NewCompanyUsecase(cfg, dependency.GetCompanyRepository(cfg)),
	}
}

// CreateCompany godoc
// @Summary Create a Company
// @Description Create a Company
// @Tags Companies
// @Accept json
// @produces json
// @Param Request body dto.CreateCompanyRequest true "Create a Company"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CompanyResponse} "Company response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/companies/ [post]
// @Security AuthBearer
func (h *CompanyHandler) Create(c *gin.Context) {
	Create(c, dto.ToCreateCompany, dto.ToCompanyResponse, h.usecase.Create)
}

// UpdateCompany godoc
// @Summary Update a Company
// @Description Update a Company
// @Tags Companies
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCompanyRequest true "Update a Company"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CompanyResponse} "Company response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/companies/{id} [put]
// @Security AuthBearer
func (h *CompanyHandler) Update(c *gin.Context) {
	Update(c, dto.ToUpdateCompany, dto.ToCompanyResponse, h.usecase.Update)
}

// DeleteCompany godoc
// @Summary Delete a Company
// @Description Delete a Company
// @Tags Companies
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse "response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/companies/{id} [delete]
// @Security AuthBearer
func (h *CompanyHandler) Delete(c *gin.Context) {
	Delete(c, h.usecase.Delete)
}

// GetCompany godoc
// @Summary Get a Company
// @Description Get a Company
// @Tags Companies
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CompanyResponse} "Company response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/companies/{id} [get]
// @Security AuthBearer
func (h *CompanyHandler) GetById(c *gin.Context) {
	GetById(c, dto.ToCompanyResponse, h.usecase.GetById)
}

// GetCompanies godoc
// @Summary Get Companies
// @Description Get Companies
// @Tags Companies
// @Accept json
// @produces json
// @Param Request body filter.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=filter.PagedList[dto.CompanyResponse]} "Company response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/companies/get-by-filter [post]
// @Security AuthBearer
func (h *CompanyHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, dto.ToCompanyResponse, h.usecase.GetByFilter)
}
