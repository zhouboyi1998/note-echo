package router

import (
	"github.com/labstack/echo/v4"
	"note-echo/src/controller"
)

// RegisterRouter 注册路由
func RegisterRouter(e *echo.Echo) {
	// 新建路由组
	g := e.Group("/note/command")
	{
		g.GET("/list", controller.List)
	}
}
