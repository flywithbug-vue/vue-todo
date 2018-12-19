package handler

import (
	"net/http"
	"todo-go/common"
	"todo-go/core/jwt"
	"todo-go/model"

	log "github.com/flywithbug/log4go"
	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(aRes.Code, aRes)
	}()
	user := new(model.User)
	err := c.BindJSON(user)
	if err != nil {
		aRes.SetErrorInfo(http.StatusBadRequest, "para invalid"+err.Error())
		return
	}
	err = user.CheckLogin(user.Account, user.Password)
	if err != nil {
		log.Error(err.Error())
		aRes.SetErrorInfo(http.StatusUnauthorized, "account or password not right")
		return
	}
	claims := jwt.NewCustomClaims(user.UserId, user.Account)
	token, err := jwt.GenerateToken(claims)
	if err != nil {
		log.Error(err.Error())
		aRes.SetErrorInfo(http.StatusUnauthorized, "token generate error"+err.Error())
		return
	}
	userAg := c.GetHeader(common.KeyUserAgent)
	_, err = model.UserLogin(user.UserId, userAg, token, c.ClientIP())
	if err != nil {
		log.Error(err.Error())
		aRes.SetErrorInfo(http.StatusUnauthorized, "token generate error"+err.Error())
		return
	}
	aRes.SetResponseDataInfo("token", token)
}

func RegisterHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(aRes.Code, aRes)
	}()
	user := new(model.User)
	err := c.BindJSON(user)
	if err != nil {
		aRes.SetErrorInfo(http.StatusBadRequest, "para invalid"+err.Error())
		return
	}
	if user.Account == "" {
		aRes.SetErrorInfo(http.StatusBadRequest, "account can not be nil")
		return
	}
	if user.Password == "" {
		aRes.SetErrorInfo(http.StatusBadRequest, "Password can not be nil")
		return
	}
	//if user.Mail == "" {
	//	aRes.SetErrorInfo(http.StatusBadRequest, "Mail can not be nil")
	//	return
	//}
	err = user.Insert()
	if err != nil {
		aRes.SetErrorInfo(http.StatusBadRequest, err.Error())
		return
	}
	aRes.AddResponseInfo("user", user)
}
