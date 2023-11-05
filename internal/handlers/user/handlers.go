package user

import (
	"log"
	"net/http"
	// "encoding/json"
)

func RegistrateUserHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("User registration handler")

	// err := json.NewEncoder(r.Body).Encode()

	// rw.Header().Set("Content-type", "application/json")

}
