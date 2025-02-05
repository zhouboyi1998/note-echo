package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"note-echo/src/repository"
)

// List 查询所有 Linux 命令
func List(ctx echo.Context) error {
	commandArray := repository.List(ctx)
	return ctx.JSON(http.StatusOK, commandArray)
}
