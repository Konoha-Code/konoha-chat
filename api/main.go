package main

import (
	"konoha/api-go/database"
	"konoha/api-go/models"
	"konoha/api-go/server"
)

func main() {
	// Start database
	database.StartDB()

	// Create table Message
	database.GetDatabase().AutoMigrate(&models.Message{})

	// Create table User
	database.GetDatabase().AutoMigrate(&models.User{})

	// Create server
	server := server.NewServer()

	// Start server
	server.Run()
}
