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

func graceFullShutdown(ctx context.Context, ch chan os.Signal, db db.Storage, server *server.MailSystemServer) {
	for {
		select {
		case <-ch:
			close(ch)
			server.Shutdown(ctx)
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

	handlers := handlers.NewMailHandler(db)
	mux := handlers.LoadHandlers()

	serverHost := os.Getenv("SERVER_HOST")
	serverPort := os.Getenv("SERVER_PORT")
	server, err := server.NewServer(serverHost, serverPort, mux)

	if server == nil || err != nil {
		log.Fatal("Server have not started")
	}

	exit := make(chan os.Signal, 1)
	go graceFullShutdown(context, exit, db, server)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	server.Start()
}
