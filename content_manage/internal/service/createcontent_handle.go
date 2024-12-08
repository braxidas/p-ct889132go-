package service

import (
	"content_manage/api/operate"
	"content_manage/internal/biz"
	"context"
	"time"
)

func (a *AppService) CreateContent(ctx context.Context,
	req *operate.CreateContentReq) (*operate.CreateContentRsp, error) {
	content := req.Content
	uc := a.uc
	err := uc.CreateContent(ctx, &biz.Content{
		Title:          content.GetTitle(),
		VideoUrl:       content.GetVideoUrl(),
		Author:         content.GetAuthor(),
		Description:    content.GetDescription(),
		Thumbnail:      content.GetThumbnail(),
		Category:       content.GetCategory(),
		Duration:       time.Duration(content.GetDuration()),
		Resolution:     content.GetResolution(),
		Filesize:       content.GetFilesize(),
		Format:         content.GetFormat(),
		Quality:        content.GetQuality(),
		ApprovalStatus: content.GetApprovalStatus(),
	})
	if err != nil {
		return nil, err
	}
	return &operate.CreateContentRsp{}, nil
}
