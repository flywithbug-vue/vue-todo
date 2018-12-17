package handler

import (
	"net/http"
	"strconv"
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
		aRes.SetErrorInfo(http.StatusNotFound, err.Error())
		return
	}
	aRes.SetSuccessInfo(http.StatusOK, "success")
}

func GetTodoHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(http.StatusOK, aRes)
	}()

	parId := c.Param("id")
	if parId == "" {
		aRes.SetErrorInfo(http.StatusBadRequest, "para invalid")
		return
	}
	id, err := strconv.ParseInt(parId, 10, 64)
	if err != nil {
		aRes.SetErrorInfo(http.StatusBadRequest, err.Error())
		return
	}
	todo, err := model.FindTodoById(id)
	if err != nil {
		aRes.SetErrorInfo(http.StatusNotFound, err.Error())
		return
	}
	aRes.AddResponseInfo("todo", todo)
}
