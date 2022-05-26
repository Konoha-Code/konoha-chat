package routes

import (
	"konoha-chat/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(route *gin.Engine) *gin.Engine {
	main := route.Group("konoha-chat/api")

	main.GET("/ping", controllers.Ping)

	return route
}
