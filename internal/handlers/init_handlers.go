package handlers

import (
	"context"
	db "mail_system/internal/db/postgres"

	"github.com/gorilla/mux"
)

type MailHandlers struct {
	Context context.Context
	Db      *db.PostgresDB
}

func NewMailHandler(ctx context.Context, db *db.PostgresDB) *MailHandlers {
	return &MailHandlers{
		Context: ctx,
		Db:      db,
	}
}

func (mh *MailHandlers) LoadHandlers() *mux.Router {
	mux := mux.NewRouter()

	// Adding handlers
	mux.HandleFunc("/registrate/client", mh.RegistrateClient).Methods("POST")
	mux.HandleFunc("/registrate/employee", mh.RegistrateEmployee).Methods("POST")
	mux.HandleFunc("/user/{id}", mh.GetUserHandler).Methods("GET")
	mux.HandleFunc("/address", mh.CreateAddress).Methods("POST")
	mux.HandleFunc("/auth/client", mh.AuthClient).Methods("POST")

	return mux
}
