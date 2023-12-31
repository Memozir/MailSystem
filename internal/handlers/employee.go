package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"mail_system/internal/config"
	"mail_system/internal/model"
	"net/http"
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
	} else {
		user, err := handler.Db.CreateUser(
			r.Context(),
			emp.User.FirstName,
			emp.User.SecondName,
			emp.User.MiddleName,
			emp.User.Login,
			emp.User.Pass,
			emp.User.BirthDate)
		if err != nil {
			log.Printf("CREATE USER ERROR: %s", user.Err.Error())
			rw.WriteHeader(http.StatusBadRequest)
		} else {
			role, err := handler.Db.GetRoleByName(r.Context(), emp.Role.Name)
			if err != nil {
				log.Printf("GET USER ROLE ERROR: %s", user.Err.Error())
				rw.WriteHeader(http.StatusBadRequest)
			} else {
				creatorEmployee, err := handler.Db.GetEmployeeByLogin(r.Context(), emp.CreatorLogin)
				if err != nil {
					log.Printf("GET USER ROLE ERROR: %s", user.Err.Error())
					rw.WriteHeader(http.StatusBadRequest)
				} else {
					_, err := handler.Db.CreateEmployee(
						r.Context(),
						user.Val.(uint64),
						creatorEmployee.Val.(model.Employee).DepartmentId,
						role.Val.(uint8))
					if err != nil {
						log.Printf("CREATE EMPLOYEE ERROR: %s", user.Err.Error())
						rw.WriteHeader(http.StatusBadRequest)
					} else {
						log.Println("CREATE EMPLOYEE SUCCESS")
						rw.WriteHeader(http.StatusCreated)
					}
				}
			}
		}
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
			if user.Val.(model.UserAuth).RoleCode >= int8(config.AdminRole) {
				employee, err := handler.Db.GetEmployeeByLogin(r.Context(), deleteInfo.User.Login)
				if err != nil {
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
								r.Context(), deleteInfo.AddressName)
							if err != nil {
								log.Printf("DELETE ADDRESS ERROR: %s", err.Error())
								rw.WriteHeader(http.StatusBadRequest)
							} else {
								log.Printf("DELETE ADDRESS %s SUCCESS", deleteInfo.AddressName)
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

type DeleteEmployeeRequest struct {
	User          UserAuthRequest `json:"user"`
	EmployeeLogin string          `json:"employee_login"`
}

func (handler *MailHandlers) DeleteEmployee(rw http.ResponseWriter, r *http.Request) {
	var request DeleteEmployeeRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Printf("DELETE EMPLOYEE DECODE ERROR: %s", err)
		rw.WriteHeader(http.StatusBadRequest)
	} else {
		employeeDelete, err := handler.Db.GetEmployeeByLogin(r.Context(), request.EmployeeLogin)
		if err != nil {
			log.Printf("GET EMPLOYEE ID ERROR: %s", err)
			rw.WriteHeader(http.StatusBadRequest)
		} else {
			admin, err := handler.Db.GetEmployeeByLogin(r.Context(), request.EmployeeLogin)
			if err != nil {
				log.Printf("GET ADMIN ID ERROR: %s", err)
				rw.WriteHeader(http.StatusBadRequest)
			} else {
				if employeeDelete.Val.(model.Employee).DepartmentId == admin.Val.(model.Employee).DepartmentId {
					err = handler.Db.DeleteEmployee(r.Context(), employeeDelete.Val.(model.Employee).EmployeeId)
					if err != nil {
						log.Printf("DELETE EMPLOYEE ERROR: %s", err)
						rw.WriteHeader(http.StatusBadRequest)
					} else {
						log.Printf("DELETE EMPLOYEE SUCCESS")
						rw.WriteHeader(http.StatusOK)
					}
				}
			}
		}
	}
}
