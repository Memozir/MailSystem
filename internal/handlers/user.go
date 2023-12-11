package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type UserJSON struct {
	Id         uint64 `json:"id"`
	Phone      string `json:"phone"`
	Login      string `json:"login"`
	Pass       string `json:"pass"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	BirthDate  string `json:"birth_date"`
}

func (handler *MailHandlers) RegistrateUserHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("User registration handler")

	var userJSON UserJSON
	err := json.NewDecoder(r.Body).Decode(&userJSON)

	if err != nil {
		log.Printf("User decode error: %s", err)
	}

	log.Println(userJSON)
	contextCreateUser, cancelCreateUser := context.WithTimeout(r.Context(), time.Second*2)
	defer cancelCreateUser()

	userId := handler.Db.CreateUser(
		contextCreateUser,
		userJSON.FirstName,
		userJSON.SecondName,
		userJSON.Login,
		userJSON.Phone,
		userJSON.Pass,
		userJSON.BirthDate)
	rw.Header().Set("Content-type", "application/json")

	fmt.Printf("User id: %d", userId)
}

func (handler *MailHandlers) GetUserHandler(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user, err := handler.Db.GetUserById(vars["id"])

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println(user)
	rw.WriteHeader(http.StatusOK)
}
