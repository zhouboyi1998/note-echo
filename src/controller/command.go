package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"note-echo/src/repository"
)

// One 根据id查询命令
func One(ctx echo.Context) error {
	command := repository.One(ctx)
	return ctx.JSON(http.StatusOK, command)
}

// List 查询命令列表
func List(ctx echo.Context) error {
	commandArray := repository.List(ctx)
	return ctx.JSON(http.StatusOK, commandArray)
}

// Insert 新增命令
func Insert(ctx echo.Context) error {
	result, commandName := repository.Insert(ctx)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"result":  result,
		"command": commandName,
	})
}

// InsertBatch 批量新增命令
func InsertBatch(ctx echo.Context) error {
	result := repository.InsertBatch(ctx)
	return ctx.JSON(http.StatusOK, result)
}

// Update 修改命令
func Update(ctx echo.Context) error {
	result := repository.Update(ctx)
	return ctx.JSON(http.StatusOK, result)
}

// UpdateBatch 批量修改命令
func UpdateBatch(ctx echo.Context) error {
	result := repository.UpdateBatch(ctx)
	return ctx.JSON(http.StatusOK, result)
}

// Delete 删除命令
func Delete(ctx echo.Context) error {
	result, objectId := repository.Delete(ctx)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"result": result,
		"_id":    objectId,
	})
}

// DeleteBatch 批量删除命令
func DeleteBatch(ctx echo.Context) error {
	result, objectIds := repository.DeleteBatch(ctx)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"result": result,
		"_ids":   objectIds,
	})
}

// Select 查询命令
func Select(ctx echo.Context) error {
	command := repository.Select(ctx)
	return ctx.JSON(http.StatusOK, command)
}

// NameList 查询命令名称列表
func NameList(ctx echo.Context) error {
	nameArray := repository.NameList(ctx)
	return ctx.JSON(http.StatusOK, nameArray)
}
