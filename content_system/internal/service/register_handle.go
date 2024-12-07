package service

import (
	"ContentSystem/internal/dao"
	"ContentSystem/internal/model"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterReq struct {
	UserID   string `json:"user_id" binding:"required"`
	Password string `json:"password" binding:"required"`
	Nickname string `json:"nickname binding:"required"`
}

type RegisterRsp struct {
	Message string `json:"message" binding:"required"`
}

func (c *CmsApp) Register(ctx *gin.Context){
	var req RegisterReq
	if err := ctx.ShouldBindJSON(&req); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	//1.密码加密
	hashedPassword, err := encryptPassword(req.Password)
	if err != nil{
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{})
	}
	//2.账号校验
	accountDao := dao.NewAccountDao(c.db)
	isExist, err := accountDao.IsExist(req.UserID)
	if err != nil{
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	if isExist{
		ctx.JSON(http.StatusBadRequest, gin.H{"error":"账号已存在"})
	}

	//3.账号信息持久化
	err = accountDao.Create(model.Account{
		UserID: req.UserID,
		Password: hashedPassword,
		Nickname: req.Nickname,
		Ct: time.Time{},
		Ut: time.Time{},
	})
	if err != nil{
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":0,
		"msg":"ok",
		"data":&RegisterRsp{
			Message: fmt.Sprintf("注册成功"),
		},
	})
}


func encryptPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil{
		fmt.Println("bcrypt generate error = %+v", err)
		return "", err
	}
	return string(hashedPassword), nil
}