package service

import (
	"ContentSystem/internal/dao"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *CmsApp) ContentDelete(ctx *gin.Context) {
	var req ContentCreateReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	contentDao := dao.NewContentDao(c.db)
	ok, err := contentDao.IsExist(req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "内容不存在"})
		return
	}
	err = contentDao.Delete(req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": &ContentCreateRsp{
			Message: fmt.Sprintf("ok"),
		},
	})

}