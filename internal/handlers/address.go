package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type AddressRequest struct {
	Name string `json:"name"`
}

func (handler *MailHandlers) CreateAddressHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("Address registration handler")

	var request AddressRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		log.Printf("Address decode error: %s", err)
	}

	log.Println(request)
	contextCreateAddress, cancelCreateAddress := context.WithTimeout(r.Context(), time.Second*2)
	defer cancelCreateAddress()

	err = handler.Db.CreateAddress(contextCreateAddress, request.Name)

	if err != nil {
		log.Printf("Address was not created: %s", err.Error())
	}

	rw.Header().Set("Content-type", "application/json")
}
