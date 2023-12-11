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

func (handler *MailHandlers) RegistrateClient(rw http.ResponseWriter, r *http.Request) {
	var clientJSON ClientJSON
	err := json.NewDecoder(r.Body).Decode(&clientJSON)

	if err != nil {
		log.Printf("User decode error: %s", err)
	}

	log.Println(clientJSON)
	userId := handler.Db.CreateUser(
		clientJSON.Login,
		clientJSON.FirstName,
		clientJSON.SecondName,
		clientJSON.Phone,
		clientJSON.Pass,
		clientJSON.BirthDate)

	context, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	err = handler.Db.CreateClient(context, userId, clientJSON.Address)

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

func (handler *MailHandlers) AuthClient(rw http.ResponseWriter, r *http.Request) {
	var userAuth UserAuth
	err := json.NewDecoder(r.Body).Decode(&userAuth)

	if err != nil {
		log.Println(err.Error())
		rw.WriteHeader(http.StatusBadRequest)
	}

	context, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	exists, err := handler.Db.AuthUser(context, userAuth.Phone)

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
