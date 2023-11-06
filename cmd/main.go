package main

import (
	"context"
	"log"
	"mail_system/internal/db"
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
	db := db.NewDb(context)

	serverHost := os.Getenv("SERVER_HOST")
	serverPort := os.Getenv("SERVER_PORT")
	server, err := server.NewServer(serverHost, serverPort, db)

	if server == nil || err != nil {
		log.Fatal("Server have not started")
	}

	mailHandler := handlers.NewMailHandler(db)
	mux := mailHandler.LoadHandlers()

	exit := make(chan os.Signal, 1)
	go graceFullShutdown(exit, db)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	server.Start(mux)
}
