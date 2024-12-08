package service

import (
	"content_system/internal/dao"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type LoginReq struct {
	UserID   string `json:"user_id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRsp struct {
	SessionID string `json:"session_id"`
	UserID    string `json:"user_id"`
	Nickname  string `json:"nickname"`
}

func (c *CmsApp) Login(ctx *gin.Context) {
	var req LoginReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//1.账号校验
	accountDao := dao.NewAccountDao(c.db)
	account, err := accountDao.FirstByUserID(req.UserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "账号不存在"})
	}
	//2.密码鉴权
	if err = bcrypt.CompareHashAndPassword(
		[]byte(account.Password),
		[]byte(req.Password),
	); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "密码不正确"})
	}
	//3.生成sessionid存入redis
	sessionID, err := c.generateSessionID(context.Background(), req.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "系统错误"})
		return
	}
	//4.回包

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": &LoginRsp{
			SessionID: sessionID,
			UserID:    account.UserID,
			Nickname:  account.Nickname,
		},
	})
}

func (c *CmsApp) generateSessionID(ctx context.Context, userID string) (string, error) {
	sessionID := uuid.New().String()
	//key:session_id:{user_id} val:session_id 20s

	//1.session id生成
	//2.session id持久化
	sessionKey := fmt.Sprintf("session_id:%s", userID)
	err := c.rdb.Set(ctx, sessionKey, sessionID, time.Hour*8).Err()
	if err != nil {
		fmt.Printf("rdb set error = %v \n", err)
		return sessionID, err
	}
	//后续鉴权使用
	authKey := fmt.Sprintf("auth_session_id:%s", sessionID)
	err = c.rdb.Set(ctx, authKey, time.Now().Unix(), time.Hour*8).Err()
	if err != nil {
		fmt.Sprintf("rdb set error = %v \n", err)
		return sessionID, err
	}

	return sessionID, nil
}
