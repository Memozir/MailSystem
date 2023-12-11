package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type EmployeeJSON struct {
	RoleName string `json:"role_name"`
}

func (emp EmployeeJSON) String() string {
	return fmt.Sprintf("Role name: %s", emp.RoleName)
}

func (handler *MailHandlers) RegistrateEmployee(rw http.ResponseWriter, r *http.Request) {
	var emp EmployeeJSON
	err := json.NewDecoder(r.Body).Decode(&emp)

	if err != nil {
		log.Printf("Registration employee error: %s", err.Error())
	}

	fmt.Println(emp)
}
