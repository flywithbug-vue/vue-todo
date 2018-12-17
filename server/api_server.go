package server

import (
	"fmt"
	"todo-go/server/handler"
	"todo-go/server/middleware"

	"github.com/gin-gonic/gin"
)

func StartApi(port string, rPrefix string, authPrefix string) {
	r := gin.New()
	r.Use(middleware.Logger(), gin.Recovery())
	r.Use(middleware.CookieMiddleware())
	handler.RegisterRouters(r, rPrefix, authPrefix)
	err := r.Run(port)
	if err != nil {
		panic(fmt.Errorf("server启动失败 %s", err.Error()))
	}
}
