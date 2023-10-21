package main

import (
	"context"
	"log"
	"url_shorter/db"
	"url_shorter/server"
	utils "url_shorter/utils"
)

func init() {
	utils.LoadEnv()
	utils.LoadHandlers()
}

func main() {
	db, err := db.NewPostgresDb()

	if err != nil {
		log.Default()
	}
	ctx := context.Background()

	db.CreateTables(ctx)

	server, err := server.NewServer("localhost", "8080", db)

	if err != nil {
		log.Fatal("Server have not strted")
	}

	log.Print("Server Successfuly Started")
	server.Start()
}
