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
		g.GET("/one/:commandId", controller.One)
		g.GET("/one/name/:commandName", controller.OneByName)
		g.GET("/list", controller.List)
		g.GET("/list/name", controller.ListByName)
		g.POST("/insert", controller.InsertOne)
		g.POST("/insert/many", controller.InsertMany)
		g.PUT("/update", controller.UpdateOne)
		g.DELETE("/delete/:commandId", controller.DeleteOne)
		g.DELETE("/delete/many", controller.DeleteMany)
	}
}
