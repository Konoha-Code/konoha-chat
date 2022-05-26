package main

import "konoha-chat/server"

func main() {
	server := server.NewServer()
	server.Run()
}
