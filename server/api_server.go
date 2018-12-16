package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"todo-go/server/handler"
	"todo-go/server/middleware"
)

func StartApi(port string, rPrefix string, auth_prefix string) {
	r := gin.New()
	r.Use(middleware.Logger(), gin.Recovery())
	r.Use(middleware.CookieMiddleware())
	handler.RegisterRouters(r, rPrefix, auth_prefix)
	err := r.Run(port)
	if err != nil {
		panic(fmt.Errorf("server启动失败 %s", err.Error()))
	}
}
