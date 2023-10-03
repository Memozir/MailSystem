package main

import (
	"url_shorter/server"
	utils "url_shorter/utils"
)

func init() {
	utils.LoadEnv()
	utils.LoadHandlers()
}

func main() {
	server := server.NewServer("localhost", "8080")
	server.Start()
}
