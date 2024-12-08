package data

import (
	"content_manage/internal/biz"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

//数据仓库
type contentRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewContentRepo(data *Data, logger log.Logger) biz.ContentRepo {
	return &contentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
func(c *contentRepo)Create(ctx context.Context, content *biz.Content)error{
	c.log.Infof("create repo error:%v",content)
	return nil
}