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

func AddTodoHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(http.StatusOK, aRes)
	}()
	title := c.PostForm("title")
	if title == "" {
		log4go.Info("title can not be nil")
		aRes.SetErrorInfo(http.StatusInternalServerError, "title can not be nil")
		return
	}
	todo := new(model.Todo)
	todo.Title = title
	err := model.InsertTodo(todo)
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusInternalServerError, err.Error())
		return
	}
	aRes.AddResponseInfo("todo", todo)
}
