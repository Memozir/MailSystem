package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

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
		log.Printf("Registration employee error: %s", err.Error())
	}
	contextCreateUser, cancelCreateuser := context.WithTimeout(context.Background(), time.Second * 2)

	user := handler.Db.CreateUser(contextCreateUser, emp.User.FirstName, emp.User.SecondName, emp.User.Login, emp.User.Phone, emp.User.Pass, emp.User.BirthDate)

	fmt.Println(emp)
}
