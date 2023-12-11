package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type RoleJSON struct {
	Code uint8
	Name string
}

func (handler *MailHandlers) CreateRoleHandler(rw http.ResponseWriter, r *http.Request) {
	var role RoleJSON
	json.NewDecoder(r.Body).Decode(&role)

	context, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	err := handler.Db.CreateRole(context, role.Code, role.Name)

	if err != nil {
		log.Printf("Role was not created: %s", err.Error())
		rw.WriteHeader(http.StatusBadRequest)
	} else {
		rw.WriteHeader(http.StatusCreated)
	}
}
