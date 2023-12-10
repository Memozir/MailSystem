package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type AddresJson struct {
	Name string `json:"name"`
}

func (handler *MailHandlers) CreateAddress(rw http.ResponseWriter, r *http.Request) {
	log.Println("Address registration handler")

	var addresJson AddresJson
	err := json.NewDecoder(r.Body).Decode(&addresJson)

	if err != nil {
		log.Printf("Address decode error: %s", err)
	}

	log.Println(addresJson)
	err = handler.Db.CreateAddress(handler.Context, addresJson.Name)

	if err != nil {
		log.Printf("Address was not created: %s", err.Error())
	}

	rw.Header().Set("Content-type", "application/json")
}
