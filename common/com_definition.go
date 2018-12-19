package common

import (
	"github.com/gin-gonic/gin"
)

const (
	KeyUserToken = "Access_token"
	KeyJWTClaims = "_key_jwt_Claims"
	KeyUserId    = "user_id"
	KeyUserAgent = "User-Agent"
)

func UserToken(ctx *gin.Context) string {
	token := ctx.GetHeader(KeyUserToken)
	return token
}
