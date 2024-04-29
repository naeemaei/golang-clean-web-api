package router

import (
	"github.com/gin-gonic/gin"
	"github.com/naeemaei/golang-clean-web-api/api/handler"
)

func TestRouter(r *gin.RouterGroup) {
	h := handler.NewTestHandler()

	r.GET("/", h.Test)
	r.GET("/users", h.Users)
	r.GET("/user/:id", h.UserById)
	r.GET("/user/get-user-by-username/:username", h.UserByUsername)
	r.GET("/user/:id/accounts", h.Accounts)
	r.POST("/add-user", h.AddUser)

	r.POST("/binder/header1", h.HeaderBinder1)
	r.POST("/binder/header2", h.HeaderBinder2)

	r.POST("/binder/query1", h.QueryBinder1)
	r.POST("/binder/query2", h.QueryBinder2)

	r.POST("/binder/uri/:id/:name", h.UriBinder)
	r.POST("/binder/body", h.BodyBinder)
	r.POST("/binder/form", h.FormBinder)
	r.POST("/binder/file", h.FileBinder)

}
