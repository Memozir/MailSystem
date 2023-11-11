package main

import (
	"context"
	"log"
	"mail_system/internal/db"
	postgres "mail_system/internal/db/postgres"
	"mail_system/internal/handlers"
	"mail_system/internal/server"
	"mail_system/internal/utils"
	"os"
	"os/signal"
	"syscall"
)

func graceFullShutdown(ch chan os.Signal, db db.Storage) {
	for {
		select {
		case <-ch:
			close(ch)
			db.Reset()
			log.Fatalf("APPLICATION WAS GRACEFULLY SHUTTED DOWN!")
			return
		default:
			continue
		}
	}
}

func main() {
	utils.LoadEnv()

	context := context.Background()
	db := postgres.NewDb(context)

	serverHost := os.Getenv("SERVER_HOST")
	serverPort := os.Getenv("SERVER_PORT")
	server, err := server.NewServer(serverHost, serverPort, db)

	if server == nil || err != nil {
		log.Fatal("Server have not started")
	}

	handlers := handlers.MailHandlers{
		User: postgres.User{Db: db},
	}
	mux := handlers.LoadHandlers()

	exit := make(chan os.Signal, 1)
	go graceFullShutdown(exit, db)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	server.Start(mux)
}
