package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	db "mail_system/internal/db/postgres"
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
		rw.WriteHeader(http.StatusBadRequest)
		r.Context().Done()
	}

	contextCreateUser, cancelCreateUser := context.WithTimeout(r.Context(), time.Second*2)
	defer cancelCreateUser()

	contextGetRole, cancelGetRole := context.WithTimeout(r.Context(), time.Second*2)
	defer cancelGetRole()

	userCh := make(chan db.ResultDB)
	roleCh := make(chan db.ResultDB)

	go func() {
		userCh <- handler.Db.CreateUser(
			contextCreateUser,
			cancelCreateUser,
			emp.User.FirstName,
			emp.User.SecondName,
			emp.User.Login,
			emp.User.Phone,
			emp.User.Pass,
			emp.User.BirthDate)
	}()

	go func() {
		roleCh <- handler.Db.GetRoleByName(contextGetRole, cancelGetRole, emp.Role.Name)
	}()

	var user db.ResultDB
	var role db.ResultDB

	for i := 0; i < 2; i++ {
		select {
		case user = <-userCh:
			continue
		case role = <-roleCh:
			continue
		}
	}

	contextCreateEmployee, cancelCreateEmployee := context.WithTimeout(r.Context(), time.Second*2)
	defer cancelCreateEmployee()

	//var empRes db.ResultDB
	employeeCreateResult := handler.Db.CreateEmployee(contextCreateEmployee, user.Val.(uint8), role.Val.(uint8))

	if employeeCreateResult.Err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		r.Context().Done()
	}

	rw.WriteHeader(http.StatusCreated)
}
