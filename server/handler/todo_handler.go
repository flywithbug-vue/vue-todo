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
	todo := new(model.Todo)
	err := c.BindJSON(&todo)
	if err != nil {
		aRes.SetErrorInfo(http.StatusBadRequest, "para invalid"+err.Error())
		return
	}
	if todo.Title == "" {
		aRes.SetErrorInfo(http.StatusBadRequest, "title can not be nil")
		return
	}
	err = model.InsertTodo(todo)
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusInternalServerError, err.Error())
		return
	}
	aRes.AddResponseInfo("todo", todo)
}

func DeleteTodoHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(http.StatusOK, aRes)
	}()
	todo := new(model.Todo)
	err := c.BindJSON(&todo)
	if err != nil {
		aRes.SetErrorInfo(http.StatusBadRequest, "para invalid"+err.Error())
		return
	}
	if todo.Id == 0 {
		aRes.SetErrorInfo(http.StatusBadRequest, "id can not be nil")
		return
	}
	err = todo.Remove(todo.Id)
	if err != nil {
		aRes.SetErrorInfo(http.StatusBadRequest, err.Error())
		return
	}
	aRes.SetSuccessInfo(http.StatusOK, "success")
}
