package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



func(c *CmsApp) Hello(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"message":"ok",
	})
}