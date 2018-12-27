package handler

import (
	"fmt"
	"net/http"
	"strings"
	"todo-go/common"
	"todo-go/model"

	"github.com/flywithbug/log4go"
	"github.com/gin-gonic/gin"
)

var (
	index = 0
)

// 系统状态信息
func IndexHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(http.StatusOK, aRes)
	}()
	userId, _ := c.Get(common.KeyUserId)
	log4go.Info("index handler %s", userId)

}

func NoRoute(c *gin.Context) {
	path := strings.Split(c.Request.URL.Path, "/")
	fmt.Println(path)
	if (path[1] != "") && (path[1] == "api") {
		aRes := model.NewResponse()
		aRes.Code = http.StatusNotFound
		aRes.Msg = "no route"
		c.JSON(http.StatusNotFound, aRes)
	} else {
		c.HTML(http.StatusOK, "index.html", "")
	}
}
