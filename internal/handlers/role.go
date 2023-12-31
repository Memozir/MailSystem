package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type RoleJSON struct {
	Code uint8  `json:"code"`
	Name string `json:"name"`
}

func (handler *MailHandlers) CreateRoleHandler(rw http.ResponseWriter, r *http.Request) {
	var role RoleJSON
	err := json.NewDecoder(r.Body).Decode(&role)

	if err != nil {
		log.Printf("Role was not created: %s", err.Error())
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	contextCreateRole, cancel := context.WithTimeout(r.Context(), time.Second*2)
	defer cancel()
	res, err := handler.Db.CreateRole(contextCreateRole, role.Code, role.Name)

	if err != nil {
		log.Printf("Role was not created: %s", res.Err.Error())
		rw.WriteHeader(http.StatusBadRequest)
		return
	} else {
		rw.WriteHeader(http.StatusCreated)
	}
}
