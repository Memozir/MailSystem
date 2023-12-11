package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type RoleJSON struct {
	Code uint8  `json:"code"`
	Name string `json:"name"`
}

type EmployeeJSON struct {
	User UserJSON `json:"user"`
	Role RoleJSON `json:"role"`
}

func (emp EmployeeJSON) String() string {
	return fmt.Sprintf("UserId: %d, Role: %d", emp.User.Id, emp.Role.Code)
}

func (handler *MailHandlers) RegistrateEmployee(rw http.ResponseWriter, r *http.Request) {
	var emp EmployeeJSON
	err := json.NewDecoder(r.Body).Decode(&emp)

	if err != nil {
		log.Panicf("Registration employee error: %s", err.Error())
	}

	fmt.Println(emp)
}
