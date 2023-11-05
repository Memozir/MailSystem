package main

import (
	"context"
	"log"
	"mail_system/internal/db"
	"mail_system/internal/handlers"
	"mail_system/internal/server"
	utils "mail_system/internal/utils"
	"os"
)

func main() {
	utils.LoadEnv()

	context := context.Background()
	db := db.NewDb(context)

	db.CreateUser("Yura", "Regan", "89100946599", "qwerty", "2002-02-15")

	serverHost := os.Getenv("SERVER_HOST")
	serverPort := os.Getenv("SERVER_PORT")
	server, err := server.NewServer(serverHost, serverPort, db)

	if err != nil {
		log.Fatal("Server have not strted")
	}

	mux := handlers.LoadHandlers()
	server.Start(mux)
}
