package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rigoncs/TodoList/application/user"
	"github.com/rigoncs/TodoList/interfaces/types"
	"net/http"
)

func UserRegisterHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserReq
		err := ctx.ShouldBind(&req)
		if err != nil {
			ctx.JSON(http.StatusOK, types.RespError(err, "user register binding failed"))
			return
		}
		userEntity := types.UserReq2Entity(&req)
		resp, err := user.ServiceImplInst.Register(ctx, userEntity)
		if err != nil {
			ctx.JSON(http.StatusOK, types.RespError(err, "user register failed"))
			return
		}
		ctx.JSON(http.StatusOK, types.RespSuccessWithData(resp))
	}
}

func UserLoginHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserReq
		err := ctx.ShouldBind(&req)
		if err != nil {
			ctx.JSON(http.StatusOK, types.RespError(err, "bind req"))
			return
		}
		entity := types.UserReq2Entity(&req)
		resp, err := user.ServiceImplInst.Login(ctx, entity)
		if err != nil {
			ctx.JSON(http.StatusOK, types.RespError(err, "user login failed"))
			return
		}
		ctx.JSON(http.StatusOK, types.RespSuccessWithData(resp))
	}
}
