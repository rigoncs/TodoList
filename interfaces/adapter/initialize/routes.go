package initialize

import (
	"github.com/gin-gonic/gin"
	api "github.com/rigoncs/TodoList/interfaces/controller"
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
		v1.POST("user/login", api.UserLoginHandler())
		v1.POST("user/register", api.UserRegisterHandler())
		// 备忘录操作部分
	}
	return r
}
