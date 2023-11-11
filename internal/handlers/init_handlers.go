package handlers

import (
	db "mail_system/internal/db/postgres"
	postgres "mail_system/internal/db/postgres"

	"github.com/gorilla/mux"
)

type MailHandlers struct {
	User postgres.User
}

func NewMailHandler(db *db.PostgresDB) *MailHandlers {
	return &MailHandlers{
		User: postgres.User{Db: db}}
}

func (mh *MailHandlers) LoadHandlers() *mux.Router {
	mux := mux.NewRouter()

	// Adding handlers
	mux.HandleFunc("/registration", mh.RegistrateUserHandler).Methods("POST")
	mux.HandleFunc("/user/{id}", mh.GetUserHandler).Methods("GET")

	return mux
}
