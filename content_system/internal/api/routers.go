package api

import (
	"content_system/internal/service"

	"github.com/gin-gonic/gin"
)

const (
	rootPath    = "/api/"
	notAuthPath = "/out/api"
)

func CmsRouters(r *gin.Engine) {
	cmsApp := service.NewCmsApp()
	session := &SessionAuth{}

	root := r.Group(rootPath).Use(session.Auth)
	{
		root.GET("/hello", cmsApp.Hello)
		root.POST("cms/content/create", cmsApp.ContentCreate)
		root.POST("cms/content/update", cmsApp.ContentUpdate)
		root.POST("cms/content/delete", cmsApp.ContentDelete)

	}

	noAuth := r.Group(notAuthPath).Use(session.Auth)
	{
		noAuth.POST("/cms/login", cmsApp.Login)
		noAuth.POST("/cms/register", cmsApp.Register)
	}
}
