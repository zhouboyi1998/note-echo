package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"note-echo/src/application"
	"note-echo/src/router"
)

func main() {
	// 新建 Echo 实例
	app := echo.New()

	// 加载日志中间件 (记录请求日志)
	app.Use(middleware.Logger())
	// 加载崩溃恢复中间件 (捕获程序中的 panic, 并恢复程序的正常运行, 防止程序因未处理的 panic 而崩溃)
	app.Use(middleware.Recover())

	// 注册路由
	router.RegisterRouter(app)
	// 启动服务
	app.Logger.Fatal(app.Start(application.App.Server.Host + ":" + application.App.Server.Port))
}
