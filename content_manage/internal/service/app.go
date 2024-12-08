package service

import (
	"content_manage/api/operate"
	"content_manage/internal/biz"
)

type AppService struct{
	//继承
	operate.UnimplementedAppServer

	uc *biz.ContentUsecase
}

func NewAppService(uc *biz.ContentUsecase)*AppService{
	return &AppService{uc:uc}
}