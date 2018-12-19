package middleware

import (
	"net/http"
	"todo-go/common"
	"todo-go/core/jwt"
	"todo-go/model"

	"github.com/gin-gonic/gin"
)

//JWTAuthMiddleware
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		aRes := model.NewResponse()
		token := c.GetHeader(common.KeyUserToken)
		if token == "" {
			aRes.SetErrorInfo(http.StatusUnauthorized, "请求未携带token，无权限访问")
			c.JSON(http.StatusUnauthorized, aRes)
			c.Abort()
			return
		}
		l, err := model.FindLoginByToken(token)
		if err != nil {
			aRes.SetErrorInfo(http.StatusUnauthorized, "未查询到Token，无权限访问")
			c.JSON(http.StatusUnauthorized, aRes)
			c.Abort()
			return
		}
		if l.Status != 1 {
			aRes.SetErrorInfo(http.StatusUnauthorized, "授权已失效")
			c.JSON(http.StatusUnauthorized, aRes)
			c.Abort()
			return
		}
		claims, err := jwt.ParseToken(token)
		if err != nil {
			model.UpdateLoginStatus(token, 0)
			if err == jwt.TokenExpired {
				aRes.SetErrorInfo(http.StatusUnauthorized, "授权已过期")
				c.JSON(http.StatusUnauthorized, aRes)
				c.Abort()
				return
			}
			aRes.SetErrorInfo(http.StatusUnauthorized, err.Error())
			c.JSON(http.StatusUnauthorized, aRes)
			c.Abort()
			return
		}
		c.Set(common.KeyUserToken, token)
		c.Set(common.KeyUserId, claims.UserId)
		c.Set(common.KeyJWTClaims, claims)
	}
}
