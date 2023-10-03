package init_utils

import (
	"net/http"

	"github.com/joho/godotenv"

	"url_shorter/handlers"
)

func LoadEnv() {
	godotenv.Load()
}

func LoadHandlers() *http.ServeMux {
	mux := http.NewServeMux()

	// Adding handlers
	mux.Handle("/", http.HandlerFunc(handlers.IndexHandler))

	return mux
}
