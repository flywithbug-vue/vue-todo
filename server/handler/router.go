package handler

import (
	"strings"
	"todo-go/server/middleware"

	"github.com/gin-gonic/gin"
)

type State int

const (
	routerTypeNormal State = iota
	routerTypeNeedAuth
)

type ginHandleFunc struct {
	handler    gin.HandlerFunc
	routerType State
	method     string
	route      string
}

//host:port/auth_prefix/prefix/path
func RegisterRouters(r *gin.Engine, prefix string, authPrefix string) {
	jwtR := r.Group(prefix + authPrefix)
	jwtR.Use(middleware.JWTAuthMiddleware())
	for _, v := range routers {
		route := strings.ToLower(v.route)
		switch v.routerType {
		case routerTypeNeedAuth:
			route = strings.ToLower(prefix + v.route)
			funcDoRouteNeedAuthRegister(strings.ToUpper(v.method), strings.ToLower(route), v.handler, jwtR)
		case routerTypeNormal:
			funcDoRouteRegister(strings.ToUpper(v.method), strings.ToLower(route), v.handler, r)
		}
	}
}

//使用jwt过滤的routerType==routerTypeNeedAuth
func funcDoRouteNeedAuthRegister(method, route string, handler gin.HandlerFunc, jwtR *gin.RouterGroup) {
	switch method {
	case "POST":
		jwtR.POST(route, handler)
	case "PUT":
		jwtR.PUT(route, handler)
	case "HEAD":
		jwtR.HEAD(route, handler)
	case "DELETE":
		jwtR.DELETE(route, handler)
	case "OPTIONS":
		jwtR.OPTIONS(route, handler)
	default:
		jwtR.GET(route, handler)
	}
}

//普通routerType==routerTypeNormal
func funcDoRouteRegister(method, route string, handler gin.HandlerFunc, r *gin.Engine) {
	switch method {
	case "POST":
		r.POST(route, handler)
	case "PUT":
		r.PUT(route, handler)
	case "HEAD":
		r.HEAD(route, handler)
	case "DELETE":
		r.DELETE(route, handler)
	case "OPTIONS":
		r.OPTIONS(route, handler)
	default:
		r.GET(route, handler)
	}
}

var routers = []ginHandleFunc{
	{
		handler:    IndexHandler,
		routerType: routerTypeNormal,
		method:     "GET",
		route:      "/",
	},
	{
		handler:    LoginHandler,
		routerType: routerTypeNormal,
		method:     "POST",
		route:      "/login",
	},
	{
		handler:    TodoListHandler,
		routerType: routerTypeNormal,
		method:     "GET",
		route:      "/todo/list",
	},
	{
		handler:    AddTodoHandler,
		routerType: routerTypeNormal,
		method:     "POST",
		route:      "/todo/add",
	},
	{
		handler:    DeleteTodoHandler,
		routerType: routerTypeNormal,
		method:     "POST",
		route:      "/todo/delete/:id",
	},
	{
		handler:    GetTodoHandler,
		routerType: routerTypeNormal,
		method:     "GET",
		route:      "/todo/item/:id",
	},
	{
		handler:    UpdateTodoHandler,
		routerType: routerTypeNormal,
		method:     "POST",
		route:      "/todo/item",
	},
	{
		handler:    CheckAllTodoHandler,
		routerType: routerTypeNormal,
		method:     "POST",
		route:      "/todo/check",
	},
	{
		handler:    DestroyCompletedItemsHandler,
		routerType: routerTypeNormal,
		method:     "POST",
		route:      "/todo/destroy",
	},
}
