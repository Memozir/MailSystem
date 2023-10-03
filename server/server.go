package server

import (
	"fmt"
	"net/http"

	"url_shorter/handlers"
)

type Server struct {
	host string
	port string
}

func NewServer(host string, port string) (server *Server) {
	// TODO: Add error returning
	// if len(host) == 0 {
	// 	return nil, "Given port is empty"
	// }

	// if len(port) == 0 {
	// 	return nil, "Given port is empty"
	// }
	server = &Server{host, port}

	return server
}

func (server *Server) Start() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(handlers.IndexHandler))
	path := fmt.Sprintf("%s:%s", server.host, server.port)
	http.ListenAndServe(path, mux)
}
