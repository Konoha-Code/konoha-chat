package main

import "konoha-chat/server"

func main() {
	// Start database
	database.StartDB()
	server := server.NewServer()
	server.Run()
}
