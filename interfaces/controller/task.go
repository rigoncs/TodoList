package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rigoncs/TodoList/application/task"
	"github.com/rigoncs/TodoList/infrastructure/common/log"
	"github.com/rigoncs/TodoList/interfaces/types"
	"net/http"
)

func CreateTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req *types.CreateTaskReq
		err := ctx.ShouldBind(&req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, types.RespError(err, "invalid request"))
			return
		}
		result, err := task.ServiceImplInst.CreateTask(ctx.Request.Context(), req)
		if err != nil {
			ctx.JSON(http.StatusOK, types.RespError(err, "create task failed"))
			return
		}
		resp := types.Entity2TaskResp(result)
		ctx.JSON(http.StatusOK, types.RespSuccessWithData(resp))
	}
}

func ListTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ListTasksReq
		err := ctx.ShouldBind(&req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, types.RespError(err, "invalid request"))
			return
		}
		resp, err := task.ServiceImplInst.ListTask(ctx.Request.Context(), &req)
		if err != nil {
			ctx.JSON(http.StatusOK, types.RespError(err, "list task failed"))
			return
		}
		ctx.JSON(http.StatusOK, types.RespSuccessWithData(resp))
	}
}

func DetailTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.DetailReq
		err := ctx.ShouldBind(&req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, types.RespError(err, "invalid request"))
			return
		}
		result, err := task.ServiceImplInst.DetailTask(ctx.Request.Context(), &req)
		if err != nil {
			ctx.JSON(http.StatusOK, types.RespError(err, "detail task failed"))
			return
		}
		resp := types.Entity2TaskResp(result)
		ctx.JSON(http.StatusOK, types.RespSuccessWithData(resp))
	}
}

func DeleteTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.DeleteTaskReq
		err := ctx.ShouldBind(&req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, types.RespError(err, "invalid request"))
			return
		}
		err = task.ServiceImplInst.DeleteTask(ctx.Request.Context(), &req)
		if err != nil {
			ctx.JSON(http.StatusOK, types.RespError(err, "delete task failed"))
			return
		}
		ctx.JSON(http.StatusOK, types.RespSuccess())
	}
}

func UpdateTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UpdateTaskReq
		err := ctx.ShouldBind(&req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, types.RespError(err, "invalid request"))
			return
		}
		err = task.ServiceImplInst.UpdateTask(ctx.Request.Context(), &req)
		if err != nil {
			ctx.JSON(http.StatusOK, types.RespError(err, "update task failed"))
			return
		}
		ctx.JSON(http.StatusOK, types.RespSuccess())
	}
}

func SearchTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.SearchTaskReq
		err := ctx.ShouldBind(&req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, types.RespError(err, "invalid request"))
			return
		}
		resp, err := task.ServiceImplInst.SearchTask(ctx.Request.Context(), &req)
		if err != nil {
			ctx.JSON(http.StatusOK, types.RespError(err, "failed to search task"))
			return
		}
		ctx.JSON(http.StatusOK, types.RespSuccessWithData(resp))
	}
}
