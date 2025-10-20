package initialize

import (
	"github.com/gin-gonic/gin"
	api "github.com/rigoncs/TodoList/interfaces/controller"
	"github.com/rigoncs/TodoList/interfaces/middleware"
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
		authed := v1.Group("/task/")
		authed.Use(middleware.JWT())
		{
			authed.POST("create", api.CreateTaskHandler())
			authed.GET("list", api.ListTaskHandler())
			authed.POST("update", api.UpdateTaskHandler())
			authed.GET("detail", api.DetailTaskHandler())
			authed.POST("search", api.SearchTaskHandler())
			authed.POST("delete", api.DeleteTaskHandler())
		}
	}
	return r
}
