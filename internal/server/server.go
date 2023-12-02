package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type MailSystemServer struct {
	httpServer *http.Server
}

func NewServer(host string, port string, router *mux.Router) (server *MailSystemServer, err error) {
	// TODO: Add error returning
	if len(host) == 0 {
		return nil, err
	}

	if len(port) == 0 {
		return nil, err
	}

	path := fmt.Sprintf("%s:%s", host, port)
	server = &MailSystemServer{
		httpServer: &http.Server{Addr: path, Handler: router},
	}

	return server, err
}

func (server *MailSystemServer) Start() {
	log.Printf("Server is starting  on %s...", server.httpServer.Addr)
	log.Fatal(server.httpServer.ListenAndServe())
}

func (server *MailSystemServer) Shutdown(ctx context.Context) {
	log.Printf("Server shutdown")
	server.httpServer.Shutdown(ctx)
}
