package server

import (
	"fmt"
	"net/http"

	"url_shorter/internal/db"
	"url_shorter/internal/handlers"
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

func (server *Server) Start() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(handlers.IndexHandler))
	path := fmt.Sprintf("%s:%s", server.host, server.port)
	http.ListenAndServe(path, mux)
}
