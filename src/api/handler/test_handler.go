package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naeemaei/golang-clean-web-api/api/helper"
)

type header struct {
	UserId  string
	Browser string
}

type personData struct {
	FirstName    string `json:"first_name" binding:"required,alpha,min=4,max=10"`
	LastName     string `json:"last_name" binding:"required,alpha,min=6,max=20"`
	MobileNumber string `json:"mobile_number" binding:"required,mobile,min=11,max=11"`
}
type TestHandler struct {
}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (h *TestHandler) Test(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"result": "Test",
	})
	c.JSON(http.StatusOK, helper.GenerateBaseResponse("Test", true, 0))

}

func (h *TestHandler) Users(c *gin.Context) {

	c.JSON(http.StatusOK, helper.GenerateBaseResponse("Users", true, 0))

}

// UserById godoc
// @Summary UserById
// @Description UserById
// @Tags Test
// @Accept  json
// @Produce  json
// @Param id path int true "user id"
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/test/user/{id} [get]
func (h *TestHandler) UserById(c *gin.Context) {

	id := c.Param("id")
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "UserById",
		"id":     id,
	}, true, 0))

}

func (h *TestHandler) UserByUsername(c *gin.Context) {

	username := c.Param("username")
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result":   "UserByUsername",
		"username": username,
	}, true, 0))
}

func (h *TestHandler) Accounts(c *gin.Context) {

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "Accounts",
	}, true, 0))
}

func (h *TestHandler) AddUser(c *gin.Context) {

	id := c.Param("id")
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "AddUser",
		"id":     id,
	}, true, 0))
}

func (h *TestHandler) HeaderBinder1(c *gin.Context) {
	userId := c.GetHeader("UserId")

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "HeaderBinder1",
		"userId": userId,
	}, true, 0))

}

func (h *TestHandler) HeaderBinder2(c *gin.Context) {
	header := header{}
	_ = c.BindHeader(&header)
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "HeaderBinder1",
		"header": header,
	}, true, 0))
}

func (h *TestHandler) QueryBinder1(c *gin.Context) {
	id := c.Query("id")
	name := c.Query("name")

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "QueryBinder1",
		"id":     id,
		"name":   name,
	}, true, 0))
}

func (h *TestHandler) QueryBinder2(c *gin.Context) {
	ids := c.QueryArray("id")
	name := c.Query("name")

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "QueryBinder2",
		"ids":    ids,
		"name":   name,
	}, true, 0))
}

// BodyBinder godoc
// @Summary BodyBinder
// @Description BodyBinder
// @Tags Test
// @Accept  json
// @Produce  json
// @Param id path int true "user id"
// @Param name path string true "user name"
// @Success 200 {object} helper.BaseHttpResponse{validationErrors=any{}} "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/test/binder/uri/{id}/{name} [post]
// @Security AuthBearer
func (h *TestHandler) UriBinder(c *gin.Context) {
	id := c.Param("id")
	name := c.Param("name")
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "UriBinder",
		"id":     id,
		"name":   name,
	}, true, 0))
}

// BodyBinder godoc
// @Summary BodyBinder
// @Description BodyBinder
// @Tags Test
// @Accept  json
// @Produce  json
// @Param person body personData true "person data"
// @Success 200 {object} helper.BaseHttpResponse{validationErrors=any{}} "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/test/binder/body [post]
// @Security AuthBearer
func (h *TestHandler) BodyBinder(c *gin.Context) {
	p := personData{}
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil,
				false, helper.ValidationError, err))
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "BodyBinder",
		"person": p,
	}, true, 0))
}

func (h *TestHandler) FormBinder(c *gin.Context) {
	p := personData{}
	_ = c.ShouldBind(&p)
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "FormBinder",
		"person": p,
	}, true, 0))
}

func (h *TestHandler) FileBinder(c *gin.Context) {
	file, _ := c.FormFile("file")
	err := c.SaveUploadedFile(file, "file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			helper.GenerateBaseResponseWithError(nil, false, helper.ValidationError, err))
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "FileBinder",
		"file":   file.Filename,
	}, true, 0))
}
