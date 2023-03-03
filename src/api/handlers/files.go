package handlers

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"

	_ "github.com/naeemaei/golang-clean-web-api/api/dto"
	_ "github.com/naeemaei/golang-clean-web-api/api/helper"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/naeemaei/golang-clean-web-api/api/dto"
	"github.com/naeemaei/golang-clean-web-api/api/helper"
	"github.com/naeemaei/golang-clean-web-api/config"
	"github.com/naeemaei/golang-clean-web-api/pkg/logging"
	"github.com/naeemaei/golang-clean-web-api/services"
)

type FileHandler struct {
	service *services.FileService
}

func NewFileHandler(cfg *config.Config) *FileHandler {
	return &FileHandler{
		service: services.NewFileService(cfg),
	}
}

// CreateFile godoc
// @Summary Create a file
// @Description Create a file
// @Tags Files
// @Accept x-www-form-urlencoded
// @produces json
// @Param file formData dto.UploadFileRequest true "Create a file"
// @Param file formData file true "Create a file"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.FileResponse} "File response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/files/ [post]
// @Security AuthBearer
func (h *FileHandler) Create(c *gin.Context){
	upload:= dto.UploadFileRequest{}
	err := c.ShouldBind(&upload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}
	req:=dto.CreateFileRequest{}
	req.Description = upload.Description
	req.MimeType = upload.File.Header.Get("Content-Type")
	req.Directory ="uploads"
	req.Name, err = saveUploadFile(upload.File, req.Directory)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	res, err := h.service.Create(c, &req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(res, true, helper.Success))


}

// UpdateFile godoc
// @Summary Update a file
// @Description Update a file
// @Tags Files
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateFileRequest true "Update a file"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.FileResponse} "File response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/files/{id} [put]
// @Security AuthBearer
func (h *FileHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

// DeleteFile godoc
// @Summary Delete a file
// @Description Delete a file
// @Tags Files
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse "response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/files/{id} [delete]
// @Security AuthBearer
func (h *FileHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	if id == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound,
			helper.GenerateBaseResponse(nil, false, helper.ValidationError))
		return
	}
	file,err := h.service.GetById(c,id)
	if err != nil {
		logger.Error(logging.IO, logging.RemoveFile, err.Error(), nil)
		c.AbortWithStatusJSON(http.StatusNotFound,
			helper.GenerateBaseResponse(nil, false, helper.NotFoundError))
		return
	}
	err = os.Remove(fmt.Sprintf("%s/%s", file.Directory, file.Name))
	if err != nil {
		logger.Error(logging.IO, logging.RemoveFile, err.Error(), nil)
		c.AbortWithStatusJSON(http.StatusNotFound,
			helper.GenerateBaseResponse(nil, false, helper.InternalError))
			return
	}
	err = h.service.Delete(c, id)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(nil, true, helper.Success))
}

// GetFile godoc
// @Summary Get a file
// @Description Get a file
// @Tags Files
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.FileResponse} "File response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/files/{id} [get]
// @Security AuthBearer
func (h *FileHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

// GetFiles godoc
// @Summary Get Files
// @Description Get Files
// @Tags Files
// @Accept json
// @produces json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.FileResponse]} "File response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/files/get-by-filter [post]
// @Security AuthBearer
func (h *FileHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFilter)
}


func saveUploadFile(file *multipart.FileHeader, directory string) (string, error){
	// test.txt -> 95239855629856.txt
	randFileName := uuid.New()
	err:= os.MkdirAll(directory,os.ModePerm)
	if err != nil{
		return "", err
	}
	fileName := file.Filename
	fileNameArr := strings.Split(fileName, ".")
	fileExt := fileNameArr[len(fileNameArr) - 1]
	fileName = fmt.Sprintf("%s.%s",randFileName,fileExt)
	dst := fmt.Sprintf("%s/%s",directory,fileName)

	src, err := file.Open()
	if err != nil{
		return "", err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil{
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	if err != nil{
		return "", err
	}
	return fileName, nil
}