package biz

import (
	// v1 "content_manage/api/helloworld/v1"
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	// "github.com/go-kratos/kratos/v2/errors"
)

var (
	// ErrUserNotFound is user not found.
)

// Greeter is a Greeter model.
type Content struct {
	Title          string    `json:"title" binding:"required"`
	VideoUrl       string    `json:"videoUrl" binding:"required"`
	Author         string    `json:"author" binding:"required"`
	Description    string    `json:"description"`
	Thumbnail      string    `json:"thumbnail"`
	Category       string    `json:"category"`
	Duration       time.Duration     `json:"duration"`
	Resolution     string    `json:"resolution"`
	Filesize       int64     `json:"filesize"`
	Format         string    `json:"format"`
	Quality        int32    `json:"quality"`
	ApprovalStatus int32     `json:"approval_status"`
	UpdateAt       time.Time `json:"update_at"`
	CreateAt       time.Time `json:"create_at"`
}

type ContentRepo interface {
	Create(context.Context, *Content) (error)
}

// GreeterUsecase is a Greeter usecase.
type ContentUsecase struct {
	repo ContentRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewContentUsecase(repo ContentRepo, logger log.Logger) *ContentUsecase {
	return &ContentUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *ContentUsecase) CreateContent(ctx context.Context, c *Content) (error) {
	uc.log.WithContext(ctx).Infof("CreateContent: %v+v", c)
	return uc.repo.Create(ctx, c)
}
