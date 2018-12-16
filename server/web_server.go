package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"todo-go/server/middleware"
)

func StartWeb(port, staticPath string) {
	r := gin.New()
	r.Use(middleware.Logger(), gin.Recovery())
	r.Static("/", staticPath)
	err := r.Run(port)
	if err != nil {
		panic(fmt.Errorf("server启动失败 %s", err.Error()))
	}
}
