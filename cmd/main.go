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

	ctx := context.Background()
	storage := postgres.NewDb(ctx)

	mailHandlers := handlers.NewMailHandler(storage)
	mux := mailHandlers.LoadHandlers()

	serverHost := os.Getenv("SERVER_HOST")
	serverPort := os.Getenv("SERVER_PORT")
	MailServer, err := server.NewServer(serverHost, serverPort, mux)

	if MailServer == nil || err != nil {
		log.Fatal("Server have not started")
	}

	exit := make(chan os.Signal, 1)
	go graceFullShutdown(ctx, exit, storage, MailServer)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	MailServer.Start()
}
