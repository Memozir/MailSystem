package handlers

import (
	"github.com/gorilla/mux"

	user "mail_system/internal/handlers/user"
)

func LoadHandlers() *mux.Router {
	mux := mux.NewRouter()

	// Adding handlers
	mux.HandleFunc("/registration", user.RegistrateUserHandler).Methods("POST")

	return mux
}
