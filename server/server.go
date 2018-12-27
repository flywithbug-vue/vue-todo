package server

import (
	"fmt"
	"time"
	"todo-go/server/handler"
	"todo-go/server/middleware"

	"github.com/gin-gonic/contrib/static"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartServer(port, staticPath, rPrefix, authPrefix string) {
	r := gin.New()
	r.Use(middleware.Logger(), gin.Recovery())
	r.Use(static.Serve("/", static.LocalFile(staticPath, true)))
	r.LoadHTMLGlob(staticPath + "/index.html")
	r.NoRoute(handler.NoRoute)
	cors.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Content-Length"},
		ExposeHeaders:    []string{},
		MaxAge:           12 * time.Hour,
		AllowCredentials: false,
	}))
	r.Use(middleware.CookieMiddleware())
	handler.RegisterRouters(r, rPrefix, authPrefix)
	err := r.Run(port)
	if err != nil {
		panic(fmt.Errorf("server启动失败 %s", err.Error()))
	}
}
