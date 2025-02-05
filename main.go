package main

import (
	"github.com/labstack/echo/v4"
	"note-echo/src/router"
)

func main() {
	// 新建 Echo 实例
	app := echo.New()
	// 注册路由
	router.RegisterRouter(app)
	// 启动服务
	app.Logger.Fatal(app.Start(":18081"))
}
