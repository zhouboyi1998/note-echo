package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	// 新建 Echo 实例
	app := echo.New()

	// Hello World
	app.GET("/hello/:name", func(c echo.Context) error {
		name := c.Param("name")
		return c.JSON(http.StatusOK, echo.Map{
			"code":    http.StatusOK,
			"message": "Hello, " + name,
		})
	})

	// 启动服务
	app.Start(":18081")
}
