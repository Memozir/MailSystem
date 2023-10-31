package init_utils

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"mail_system/internal/handlers"
)

func LoadEnv() {
	godotenv.Load()
}

func LoadHandlers() *mux.Router {
	mux := mux.NewRouter()

	// Adding handlers
	mux.Handle("/", http.HandlerFunc(handlers.IndexHandler))

	return mux
}
