package handler

import (
	"net/http"
	"todo-go/model"

	"github.com/flywithbug/log4go"
	"github.com/gin-gonic/gin"
)

func TodoListHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(http.StatusOK, aRes)
	}()
	todos, err := model.FindAllTodos()
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusInternalServerError, err.Error())
		return
	}
	aRes.AddResponseInfo("list", todos)
}
