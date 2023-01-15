package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

}

func (h *TestHandler) Users(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"result": "Users",
	})

}

func (h *TestHandler) UserById(c *gin.Context) {

	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"result": "UserById",
		"id":     id,
	})
}

func (h *TestHandler) UserByUsername(c *gin.Context) {

	username := c.Param("username")
	c.JSON(http.StatusOK, gin.H{
		"result":   "UserByUsername",
		"username": username,
	})
}

func (h *TestHandler) Accounts(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"result": "Accounts",
	})
}

func (h *TestHandler) AddUser(c *gin.Context) {

	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"result": "AddUser",
		"id":     id,
	})
}

func (h *TestHandler) HeaderBinder1(c *gin.Context) {
	userId := c.GetHeader("UserId")

	c.JSON(http.StatusOK, gin.H{
		"result": "HeaderBinder1",
		"userId": userId,
	})
}

func (h *TestHandler) HeaderBinder2(c *gin.Context) {
	header := header{}
	c.BindHeader(&header)

	c.JSON(http.StatusOK, gin.H{
		"result": "HeaderBinder1",
		"header": header,
	})
}

func (h *TestHandler) QueryBinder1(c *gin.Context) {
	id := c.Query("id")
	name := c.Query("name")

	c.JSON(http.StatusOK, gin.H{
		"result": "QueryBinder1",
		"id":     id,
		"name":   name,
	})
}

func (h *TestHandler) QueryBinder2(c *gin.Context) {
	ids := c.QueryArray("id")
	name := c.Query("name")

	c.JSON(http.StatusOK, gin.H{
		"result": "QueryBinder2",
		"ids":    ids,
		"name":   name,
	})
}

func (h *TestHandler) UriBinder(c *gin.Context) {
	id := c.Param("id")
	name := c.Param("name")

	c.JSON(http.StatusOK, gin.H{
		"result": "UriBinder",
		"id":     id,
		"name":   name,
	})
}

func (h *TestHandler) BodyBinder(c *gin.Context) {
	p := personData{}
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"validationError": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "BodyBinder",
		"person": p,
	})
}

func (h *TestHandler) FormBinder(c *gin.Context) {
	p := personData{}
	c.ShouldBind(&p)
	c.JSON(http.StatusOK, gin.H{
		"result": "FormBinder",
		"person": p,
	})
}

func (h *TestHandler) FileBinder(c *gin.Context) {
	file, _ := c.FormFile("file")
	err := c.SaveUploadedFile(file, "file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "FileBinder",
		"file":   file.Filename,
	})
}
