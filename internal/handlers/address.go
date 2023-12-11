package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type AddresJson struct {
	Name string `json:"name"`
}

func (handler *MailHandlers) CreateAddressHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("Address registration handler")

	var addresJson AddresJson
	err := json.NewDecoder(r.Body).Decode(&addresJson)

	if err != nil {
		log.Printf("Address decode error: %s", err)
	}

	log.Println(addresJson)
	contextCreateAddress, cancelCreateAddress := context.WithTimeout(r.Context(), time.Second*2)
	defer cancelCreateAddress()

	err = handler.Db.CreateAddress(contextCreateAddress, addresJson.Name)

	if err != nil {
		log.Printf("Address was not created: %s", err.Error())
	}

	rw.Header().Set("Content-type", "application/json")
}
