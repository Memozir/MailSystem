package handlers

import (
	"context"
	db "mail_system/internal/db/postgres"

	"github.com/gorilla/mux"
)

type MailHandlers struct {
	Context context.Context
	Db      db.Storage
}

func NewMailHandler(db db.Storage) *MailHandlers {
	return &MailHandlers{
		Db: db,
	}
}

func (handler *MailHandlers) LoadHandlers() *mux.Router {
	router := mux.NewRouter()

	// Adding handlers
	router.HandleFunc("/register/client", handler.RegistrateClientHandler).Methods("POST")
	router.HandleFunc("/register/employee", handler.RegistrateEmployeeHandler).Methods("POST")
	router.HandleFunc("/user/{id}", handler.GetUserHandler).Methods("GET")
	router.HandleFunc("/address", handler.CreateAddressHandler).Methods("POST")
	router.HandleFunc("/auth/client", handler.AuthClientHandler).Methods("POST")
	router.HandleFunc("/create/role", handler.CreateRoleHandler).Methods("POST")

	return router
}
