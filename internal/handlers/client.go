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
	Pass       string `json:"pass"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	MiddleName string `json:"middle_name"`
	BirthDate  string `json:"birth_date"`
	Address    string `json:"address"`
	Apartment  string `json:"apartment"`
}

type ClientCreateResponse struct {
	ClientId string `json:"id"`
}

func (handler *MailHandlers) RegisterClientHandler(rw http.ResponseWriter, r *http.Request) {
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
		cancel,
		clientJSON.FirstName,
		clientJSON.SecondName,
		clientJSON.MiddleName,
		clientJSON.Login,
		clientJSON.Pass,
		clientJSON.BirthDate)

	contextCreateClient, cancelCreateClient := context.WithTimeout(context.Background(), time.Second*2)
	defer cancelCreateClient()

	err = handler.Db.CreateClient(contextCreateClient, userId.Val.(uint8), clientJSON.Address, clientJSON.Apartment)

	if err != nil {
		log.Printf("client was not created: %s", err.Error())
		rw.WriteHeader(http.StatusBadRequest)
	} else {
		rw.WriteHeader(http.StatusCreated)
	}

	rw.Header().Set("Content-type", "application/json")
}

type UserAuth struct {
	Login string `json:"login"`
	Pass  string `json:"pass"`
}

func (handler *MailHandlers) AuthUserHandler(rw http.ResponseWriter, r *http.Request) {
	var userAuth UserAuth
	err := json.NewDecoder(r.Body).Decode(&userAuth)

	if err != nil {
		log.Println(err.Error())
		rw.WriteHeader(http.StatusBadRequest)
	}

	contextAuth, cancelAuth := context.WithTimeout(context.Background(), time.Second*2)
	defer cancelAuth()
	exists, err := handler.Db.AuthUser(contextAuth, userAuth.Login, userAuth.Pass)

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
