package main

import (
	"context"
	"log"
	"mail_system/internal/db"
	"mail_system/internal/server"
	utils "mail_system/internal/utils"
)

func init() {
	utils.LoadEnv()
	utils.LoadHandlers()
}

func main() {
	context := context.Background()
	db, err := db.NewDb(context)

	// db.CreateTables(ctx)

	server, err := server.NewServer("localhost", "8080", db)

	if err != nil {
		log.Fatal("Server have not strted")
	}

	log.Print("Server Successfuly Started")
	server.Start()
}
