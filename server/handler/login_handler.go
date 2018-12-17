package handler

import (
	"net/http"
	"todo-go/core/jwt"
	"todo-go/model"

	log "github.com/flywithbug/log4go"
	"github.com/gin-gonic/gin"
)

func LoginHandler(ctx *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		ctx.JSON(http.StatusOK, aRes)
	}()
	login := parameterLoginModel{}
	err := ctx.BindJSON(&login)
	if err != nil {
		aRes.SetErrorInfo(http.StatusBadRequest, "para invalid"+err.Error())
		return
	}
	u, err := model.CheckLogin(login.Account, login.Password)
	if err != nil {
		log.Error(err.Error())
		aRes.SetErrorInfo(http.StatusUnauthorized, "account or password not right")
		return
	}
	claims := jwt.NewCustomClaims(u.UserId, u.Account)
	token, err := jwt.GenerateToken(claims)
	if err != nil {
		log.Error(err.Error())
		aRes.SetErrorInfo(http.StatusUnauthorized, "token generate error"+err.Error())
		return
	}
	aRes.SetResponseDataInfo("token", token)
}
