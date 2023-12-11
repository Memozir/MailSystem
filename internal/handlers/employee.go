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

func (handler *MailHandlers) RegistrateEmployeeHandler(rw http.ResponseWriter, r *http.Request) {
	var emp EmployeeJSON
	err := json.NewDecoder(r.Body).Decode(&emp)

	if err != nil {
		log.Printf("Registration employee error: %s", err.Error())
	}
	contextCreateUser, cancelCreateUser := context.WithTimeout(r.Context(), time.Second*2)
	defer cancelCreateUser()

	user := handler.Db.CreateUser(
		contextCreateUser,
		emp.User.FirstName,
		emp.User.SecondName,
		emp.User.Login,
		emp.User.Phone,
		emp.User.Pass,
		emp.User.BirthDate)

	contextCreateRole, cancelCreateRole := context.WithTimeout(r.Context(), time.Second*2)
	defer cancelCreateRole()

	role, err := handler.Db.GetRoleByName(contextCreateRole, emp.Role.Name)

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	contextCreateEmployee, cancelCreateEmployee := context.WithTimeout(r.Context(), time.Second*2)
	defer cancelCreateEmployee()

	_, err = handler.Db.CreateEmployee(contextCreateEmployee, user, role)

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	fmt.Println(emp)
}
