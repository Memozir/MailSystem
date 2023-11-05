package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"mail_system/internal/db"
)

type Server struct {
	host    string
	port    string
	storage db.Storage
}

func NewServer(host string, port string, storage db.Storage) (server *Server, err error) {
	// TODO: Add error returning
	if len(host) == 0 {
		return nil, err
	}

	if len(port) == 0 {
		return nil, err
	}

	server = &Server{host, port, storage}

	return server, err
}

func (server *Server) Start(mux *mux.Router) {
	path := fmt.Sprintf("%s:%s", server.host, server.port)
	log.Printf("Server is starting  on %s...", path)
	log.Fatal(http.ListenAndServe(path, mux))
}
