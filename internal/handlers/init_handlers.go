package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"mail_system/internal/db"
)

type mailHandler struct {
	db db.Storage
	// UserHandlers
}

func NewMailHandler(db db.Storage) *mailHandler {
	return &mailHandler{db: db}
}

func (mh *mailHandler) LoadHandlers() *mux.Router {
	mux := mux.NewRouter()

	// Adding handlers
	mux.HandleFunc("/registration", mh.RegistrateUserHandler).Methods("POST")
	mux.HandleFunc("/user/{id}", mh.GetUserHandler).Methods("GET")

	return mux
}

type UserJSON struct {
	Id         uint64 `json:"id"`
	Phone      string `json:"phone"`
	Pass       string `json:"pass"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"last_name"`
	BirthDate  string `json:"birth_date"`
}

func (handler *mailHandler) RegistrateUserHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("User registration handler")

	var userJSON UserJSON
	err := json.NewDecoder(r.Body).Decode(&userJSON)

	if err != nil {
		log.Printf("User decode error: %s", err)
	}

	log.Println(userJSON)
	handler.db.CreateUser(
		userJSON.FirstName,
		userJSON.SecondName,
		userJSON.Phone,
		userJSON.Pass,
		userJSON.BirthDate)
	rw.Header().Set("Content-type", "application/json")

}

func (handler *mailHandler) GetUserHandler(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := handler.db.GetUserById(vars["id"])
	fmt.Println(user)
}
