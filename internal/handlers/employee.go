package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	db "mail_system/internal/db/postgres"
	"mail_system/internal/model"
	"net/http"
	"time"
)

type EmployeeJSON struct {
	User         UserJSON `json:"user"`
	Role         RoleJSON `json:"role"`
	CreatorLogin string   `json:"login"`
}

func (emp EmployeeJSON) String() string {
	return fmt.Sprintf("UserId: %d, Role: %d", emp.User.Id, emp.Role.Code)
}

func (handler *MailHandlers) RegisterEmployeeHandler(rw http.ResponseWriter, r *http.Request) {
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
	creatorEmployeeCh := make(chan db.ResultDB)

	go func() {
		userCh <- handler.Db.CreateUser(
			contextCreateUser,
			cancelCreateUser,
			emp.User.FirstName,
			emp.User.SecondName,
			emp.User.Login,
			emp.User.Pass,
			emp.User.MiddleName,
			emp.User.BirthDate)
	}()

	go func() {
		roleCh <- handler.Db.GetRoleByName(contextGetRole, cancelGetRole, emp.Role.Name)
	}()

	go func() {
		creatorEmployeeCh <- handler.Db.GetEmployeeByLogin(r.Context(), emp.CreatorLogin)
	}()

	var user db.ResultDB
	var role db.ResultDB
	var creatorEmployee db.ResultDB

	for i := 0; i < 2; i++ {
		select {
		case user = <-userCh:
			continue
		case role = <-roleCh:
			continue
		case creatorEmployee = <-creatorEmployeeCh:
			continue
		}
	}

	contextCreateEmployee, cancelCreateEmployee := context.WithTimeout(r.Context(), time.Second*2)
	defer cancelCreateEmployee()

	employeeCreateResult := handler.Db.CreateEmployee(
		contextCreateEmployee,
		user.Val.(uint8),
		creatorEmployee.Val.(model.Employee).DepartmentId,
		role.Val.(uint8))

	if employeeCreateResult.Err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		r.Context().Done()
	} else {
		rw.WriteHeader(http.StatusCreated)
	}
}

type ManageAddressJSON struct {
	User        UserAuthRequest `json:"user"`
	AddressName string          `json:"address"`
}

func (handler *MailHandlers) DeleteAddressByAdmin(rw http.ResponseWriter, r *http.Request) {
	var deleteInfo ManageAddressJSON

	err := json.NewDecoder(r.Body).Decode(&deleteInfo)
	if err != nil {
		log.Printf("DELETE ADDRESS ENCODE REQUEST ERROR: %s", err.Error())
		rw.WriteHeader(http.StatusBadRequest)
	} else {
		user, err := handler.Db.AuthUser(r.Context(), deleteInfo.User.Login, deleteInfo.User.Pass)
		if err != nil {
			log.Printf("DELETE ADDRESS AUTH USER ERROR: %s", err.Error())
			rw.WriteHeader(http.StatusBadRequest)
		} else {
			if user.Val.(model.UserAuth).RoleCode >= 3 {
				employee := handler.Db.GetEmployeeByLogin(r.Context(), deleteInfo.User.Login)
				if employee.Err != nil {
					log.Printf("GET ADMIN ERROR: %s", err.Error())
					rw.WriteHeader(http.StatusBadRequest)
				} else {
					check, err := handler.Db.CheckAdminAddress(
						r.Context(), employee.Val.(model.Employee).EmployeeId, deleteInfo.AddressName)
					if err != nil {
						log.Printf("CHECK ADDRESS ERROR: %s", err.Error())
						rw.WriteHeader(http.StatusBadRequest)
					} else {
						if check {
							err = handler.Db.DeleteAddress(
								r.Context(), employee.Val.(model.Employee).EmployeeId, deleteInfo.AddressName)
							if err != nil {
								log.Printf("DELETE ADDRESS ERROR: %s", err.Error())
								rw.WriteHeader(http.StatusBadRequest)
							} else {
								rw.WriteHeader(http.StatusOK)
							}
						} else {
							log.Println("NO SUCH ADDRESS IN DEPARTMENT")
							rw.WriteHeader(http.StatusBadRequest)
						}
					}
				}
			} else {
				log.Println("NOT ENOUGH RIGHTS")
				rw.WriteHeader(http.StatusBadRequest)
			}
		}
	}
}
