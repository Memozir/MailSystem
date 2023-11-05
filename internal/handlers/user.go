package handlers

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
