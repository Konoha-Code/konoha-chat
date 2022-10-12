package routes

import (
	"sysmo/api-go/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(route *gin.Engine) *gin.Engine {
	main := route.Group("konoha/chat")

	main.GET("/ping", controllers.GetPing)
	main.GET("/createUser", controllers.CreateUser)
	main.GET("/getUser", controllers.GetUser)
	main.GET("/sendMessage", controllers.SendMessage)
	main.GET("/getMessages", controllers.GetMessages)

	return route
}
