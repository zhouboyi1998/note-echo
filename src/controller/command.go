package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"note-echo/src/repository"
)

// One 按 ObjectId 查询单条 Linux 命令
func One(ctx echo.Context) error {
	command := repository.One(ctx)
	return ctx.JSON(http.StatusOK, command)
}

// OneByName 按 Linux 命令名称查询单条 Linux 命令
func OneByName(ctx echo.Context) error {
	command := repository.OneByName(ctx)
	return ctx.JSON(http.StatusOK, command)
}

// List 查询所有 Linux 命令
func List(ctx echo.Context) error {
	commandArray := repository.List(ctx)
	return ctx.JSON(http.StatusOK, commandArray)
}

// ListByName 查询所有 Linux 命令的名称
func ListByName(ctx echo.Context) error {
	nameArray := repository.ListByName(ctx)
	return ctx.JSON(http.StatusOK, nameArray)
}

// InsertOne 插入单条 Linux 命令
func InsertOne(ctx echo.Context) error {
	result, commandName := repository.InsertOne(ctx)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"result":  result,
		"command": commandName,
	})
}

// InsertMany 插入多条 Linux 命令
func InsertMany(ctx echo.Context) error {
	result := repository.InsertMany(ctx)
	return ctx.JSON(http.StatusOK, result)
}

// UpdateOne 更新单条 Linux 命令
func UpdateOne(ctx echo.Context) error {
	result := repository.UpdateOne(ctx)
	return ctx.JSON(http.StatusOK, result)
}

// DeleteOne 删除单条 Linux 命令
func DeleteOne(ctx echo.Context) error {
	result, objectId := repository.DeleteOne(ctx)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"result": result,
		"_id":    objectId,
	})
}

// DeleteMany 删除多条 Linux 命令
func DeleteMany(ctx echo.Context) error {
	result, objectIds := repository.DeleteMany(ctx)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"result": result,
		"_ids":   objectIds,
	})
}
