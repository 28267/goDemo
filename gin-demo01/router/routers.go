package router

import (
	"net/http"

	"github.com/28267/goDemo/gin-demo01/controllers"
	"github.com/28267/goDemo/gin-demo01/pkg/logger"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {

	r := gin.Default()

	r.Use(gin.LoggerWithConfig(logger.LoggerToFile()))
	r.Use(logger.Recover)

	user := r.Group("/user")
	{
		user.POST("/list", controllers.UserController{}.GetList)

		user.POST("/add", controllers.UserController{}.AddUser)
		user.POST("/update", controllers.UserController{}.UpdateUser)
		user.POST("/delete", controllers.UserController{}.DeleteUser)
		user.POST("/get", controllers.UserController{}.GetUser)

		user.GET("/hello", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "Hello,World!")
		})

		// user.POST("/list", func(ctx *gin.Context) {
		// 	ctx.String(http.StatusOK, "USR LIST")
		// })

		user.PUT("/put", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "USR PUT")
		})
	}

	order := r.Group("/order")
	{
		order.POST("/list", controllers.OrderController{}.GetList)
	}
	return r

}
