package main

import (
	"context"
	"log"
	"mail_system/internal/db"
	"mail_system/internal/server"
	utils "mail_system/internal/utils"
)

func main() {
	utils.LoadEnv()

	context := context.Background()
	db := db.NewDb(context)

	utils.LoadHandlers()

	server, err := server.NewServer("localhost", "8080", db)

	if err != nil {
		log.Fatal("Server have not strted")
	}

	log.Print("Server Successfuly Started")
	server.Start()
}
