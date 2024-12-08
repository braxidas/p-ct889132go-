package service

import (
	"content_system/internal/dao"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ContentFindReq struct {
	ID       int
	Page     int
	PageSize int
}

type Content struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ContentFindRsp struct {
	Message  string    `json:"message"`
	Contents []Content `json:"contents"`
	Total    int64     `json:"total"`
}

func (c *CmsApp) ContentFind(ctx *gin.Context) {
	var req ContentFindReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	contentDao := dao.NewContentDao(c.db)

	contentList, total, err := contentDao.Find(&dao.FindParams{
		ID:       req.ID,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {

	}
	contents := make([]Content, 0, len(contentList))
	for _, content := range contentList {
		contents = append(contents, Content{
			ID:          content.ID,
			Title:       content.Title,
			Description: content.Description,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": &ContentFindRsp{
			Message:  fmt.Sprintf("ok"),
			Contents: contents,
			Total:    total,
		},
	})
}
