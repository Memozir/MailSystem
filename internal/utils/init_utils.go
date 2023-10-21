package init_utils

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"url_shorter/internal/handlers"
)

func LoadEnv() {
	godotenv.Load()
}

func LoadHandlers() *mux.Router {
	// mux := http.NewServeMux()
	mux := mux.NewRouter()

	// Adding handlers
	mux.Handle("/", http.HandlerFunc(handlers.IndexHandler))

	return mux
}
