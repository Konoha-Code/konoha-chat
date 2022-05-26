package server

import (
	"konoha-chat/server/routes"
	"log"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	customServer := gin.Default()
	customServer.Use(cors.Default())

	return Server{
		port:   "5000",
		server: customServer,
	}
}

func (s *Server) Run() {
	route := routes.ConfigRoutes(s.server)
	log.Print("Server is running at port: " + s.port)
	log.Fatal(route.Run(":" + s.port))
}
