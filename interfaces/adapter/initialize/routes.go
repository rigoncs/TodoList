package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/rigoncs/TodoList/interfaces/controller"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, "success")
		})

		// 用户操作部分
		v1.POST("user/login", controller.UserLoginHandler())
		v1.POST("user/register", controller.UserRegisterHandler())
		// 备忘录操作部分
	}
	return r
}
