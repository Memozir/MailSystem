package main

import (
	"url_shorter/server"
)

func main() {
	server := server.NewServer("localhost", "8080")
	server.Start()
}
