package middleware

import (
	"github.com/gin-gonic/gin"
	_const "github.com/rigoncs/TodoList/const"
	"github.com/rigoncs/TodoList/infrastructure/auth"
	lctx "github.com/rigoncs/TodoList/infrastructure/common/context"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = _const.SUCCESS
		token := c.GetHeader("Authorization")
		if token == "" {
			code = http.StatusNotFound
			c.JSON(_const.InvalidParams, gin.H{
				"status": code,
				"msg":    _const.GetMsg(code),
				"data":   "缺少token",
			})
			c.Abort()
			return
		}
		jwtService := auth.NewJWTTokenService()
		claims, err := jwtService.ParseToken(c.Request.Context(), token)
		if err != nil {
			code = _const.ErrorAuthCheckTokenFail
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = _const.ErrorAuthCheckTokenTimeout
		}

		if code != _const.SUCCESS {
			c.JSON(_const.InvalidParams, gin.H{
				"status": code,
				"msg":    _const.GetMsg(code),
				"data":   "token验证失败",
			})
			c.Abort()
			return
		}

		c.Request = c.Request.WithContext(
			lctx.NewContext(
				c.Request.Context(),
				&lctx.UserInfo{Id: claims.Id, Name: claims.Username},
			),
		)
		c.Next()
	}
}
