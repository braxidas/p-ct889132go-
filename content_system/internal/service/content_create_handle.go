package service

import (
	"content_system/internal/dao"
	"content_system/internal/model"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	goflow "github.com/s8sg/goflow/v1"
)

type ContentCreateReq struct {
	ID             int    `json:"id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	VideoURL       string `json:"video_url"`
	Category       string `json:"category"`
	ApprovalStatus int    `json:"approval_status"`
	Thumbnail      string `json:"thumbnail"`
	Format         string `json:"format"`
}

type ContentCreateRsp struct {
	Message string `json:"message"`
}

func (c *CmsApp) ContentCreate(ctx *gin.Context) {
	var req ContentCreateReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	contentDao := dao.NewContentDao(c.db)
	err := contentDao.Create(model.ContentDetail{
		Title:          req.Title,
		Description:    req.Description,
		ApprovalStatus: req.ApprovalStatus,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
		return
	}

	flowData := map[string]interface{}{
		"content_id": req.ID,
	}
	data, _ := json.Marshal(flowData)
	if err := c.flowService.Execute("content-flow", &goflow.Request{
		Body: data,
	}); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
