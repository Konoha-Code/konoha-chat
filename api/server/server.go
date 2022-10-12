package server

import (
	"log"
	"sysmo/api-go/server/routes"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	s := gin.Default()
	s.Use(cors.Default())

	return Server{
		port:   "5000",
		server: s,
	}
}

func (s *Server) Run() {
	route := routes.ConfigRoutes(s.server)
	log.Print("Server is running at port: " + s.port)
	log.Fatal(route.Run(":" + s.port))
}
