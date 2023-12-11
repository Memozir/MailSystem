package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type ClientJSON struct {
	Id         uint64 `json:"id"`
	Login      string `json:"login"`
	Phone      string `json:"phone"`
	Pass       string `json:"pass"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	BirthDate  string `json:"birth_date"`
	Address    string `json:"address"`
}

type ClientCreateResponse struct {
	ClientId string `json:"id"`
}

func (handler *MailHandlers) RegistrateClientHandler(rw http.ResponseWriter, r *http.Request) {
	var clientJSON ClientJSON
	err := json.NewDecoder(r.Body).Decode(&clientJSON)

	if err != nil {
		log.Printf("User decode error: %s", err)
	}

	log.Println(clientJSON)
	contextUserCreate, cancel := context.WithTimeout(r.Context(), time.Second*2)
	defer cancel()

	userId := handler.Db.CreateUser(
		contextUserCreate,
		clientJSON.Login,
		clientJSON.FirstName,
		clientJSON.SecondName,
		clientJSON.Phone,
		clientJSON.Pass,
		clientJSON.BirthDate)

	contextCreateClient, cancelCreateClient := context.WithTimeout(context.Background(), time.Second*2)
	defer cancelCreateClient()

	err = handler.Db.CreateClient(contextCreateClient, userId, clientJSON.Address)

	if err != nil {
		log.Printf("client was not created: %s", err.Error())
		rw.WriteHeader(http.StatusBadRequest)
	} else {
		rw.WriteHeader(http.StatusCreated)
	}

	rw.Header().Set("Content-type", "application/json")
}

type UserAuth struct {
	Phone string `json:"phone"`
}

func (handler *MailHandlers) AuthClientHandler(rw http.ResponseWriter, r *http.Request) {
	var userAuth UserAuth
	err := json.NewDecoder(r.Body).Decode(&userAuth)

	if err != nil {
		log.Println(err.Error())
		rw.WriteHeader(http.StatusBadRequest)
	}

	contextAuth, cancelAuth := context.WithTimeout(context.Background(), time.Second*2)
	defer cancelAuth()
	exists, err := handler.Db.AuthUser(contextAuth, userAuth.Phone)

	if err != nil {
		log.Println(err.Error())
		rw.WriteHeader(http.StatusBadRequest)
	}

	if exists {
		log.Println("SUCCESS AUTH")
		rw.WriteHeader(http.StatusOK)
	} else {
		log.Println("ERROR AUTH")
		rw.WriteHeader(http.StatusBadRequest)
	}
}
