package utils

import (
	"fmt"

)

func GetAuthKey(sessionID string) string {
	authKey := fmt.Sprintf("auth_session_id:%s", sessionID)
	return authKey
}

func GetSessionKey(userID string)string{
	sessionKey:=fmt.Sprintf("session_id:%s",userID)
	return sessionKey
}