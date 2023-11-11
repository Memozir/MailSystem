package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type UserJSON struct {
	Id         uint64 `json:"id"`
	Phone      string `json:"phone"`
	Pass       string `json:"pass"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"last_name"`
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
	handler.User.CreateUser(
		userJSON.FirstName,
		userJSON.SecondName,
		userJSON.Phone,
		userJSON.Pass,
		userJSON.BirthDate)
	rw.Header().Set("Content-type", "application/json")

}

func (handler *MailHandlers) GetUserHandler(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := handler.User.GetUserById(vars["id"])
	fmt.Println(user)
}

// import (
// 	"encoding/json"
// 	"log"
// 	"net/http"

// 	"mail_system/internal/model"
// )

// type UserHandlers struct{}

// func (handler *UserHandlers) RegistrateUserHandler(rw http.ResponseWriter, r *http.Request) {
// 	log.Println("User registration handler")

// 	var user model.User
// 	err := json.NewDecoder(r.Body).Decode(&user)

// 	if err != nil {
// 		log.Printf("User decode error: %s", err)
// 	}

// 	log.Println(user)
// 	rw.Header().Set("Content-type", "application/json")

// }
