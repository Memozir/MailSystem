package handlers

import (
	"encoding/json"
	"log"
	"mail_system/internal/model"
	"net/http"
)

type DepartmentsResponseJSON struct {
	Departments []model.Department `json:"departments"`
}

func (handler *MailHandlers) GetAllDepartments(rw http.ResponseWriter, r *http.Request) {
	departments, err := handler.Db.GetAllDepartments(r.Context())
	if err != nil {
		log.Printf("GET ALL DEPARTMENTS ERROR: %s", err)
		rw.WriteHeader(http.StatusBadRequest)
	} else {
		response := DepartmentsResponseJSON{Departments: departments}
		err := json.NewEncoder(rw).Encode(&response)

		if err != nil {
			log.Printf("GET ALL DEPARTMENTS ENCODE RESPONSE ERROR: %s", err)
			rw.WriteHeader(http.StatusBadRequest)
		} else {
			rw.WriteHeader(http.StatusOK)
		}
	}
}

func (handler *MailHandlers) GetClientDepartments(rw http.ResponseWriter, r *http.Request) {
	var user UserAuthRequest
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Printf("GET CLIENT DEPARTMENTS ERROR: %s", err)
		rw.WriteHeader(http.StatusBadRequest)
	} else {
		clientId, err := handler.Db.GetClientByLogin(r.Context(), user.Login)

		if err != nil {
			log.Printf("GET CLENT ID ERROR: %s", err)
			rw.WriteHeader(http.StatusBadRequest)
		} else {
			departments, err := handler.Db.GetClientDepartments(r.Context(), clientId)

			if err != nil {
				log.Printf("GET CLENT DEPARTMENTS ERROR: %s", err)
				rw.WriteHeader(http.StatusBadRequest)
			} else {
				response := DepartmentsResponseJSON{Departments: departments}
				err = json.NewEncoder(rw).Encode(response)
				rw.WriteHeader(http.StatusOK)
			}
		}
	}
}

func (handler *MailHandlers) GetEmployeeDepartments(rw http.ResponseWriter, r *http.Request) {
	var user UserAuthRequest
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Printf("GET CLIENT DEPARTMENTS ERROR: %s", err)
		rw.WriteHeader(http.StatusBadRequest)
	} else {
		employee := handler.Db.GetEmployeeByLogin(r.Context(), user.Login)

		if employee.Err != nil {
			log.Printf("GET EMPLOYEE ID ERROR: %s", err)
			rw.WriteHeader(http.StatusBadRequest)
		} else {
			departments, err := handler.Db.GetEmployeeDepartments(
				r.Context(),
				employee.Val.(model.Employee).EmployeeId)

			if err != nil {
				log.Printf("GET CLENT DEPARTMENTS ERROR: %s", err)
				rw.WriteHeader(http.StatusBadRequest)
			} else {
				response := DepartmentsResponseJSON{Departments: departments}
				err = json.NewEncoder(rw).Encode(response)
				rw.WriteHeader(http.StatusOK)
			}
		}
	}
}

type GetDepartmentEmployeesResponse struct {
	Employees []model.EmployeeInfo `json:"employees"`
}

func (handler *MailHandlers) GetDepartmentEmployees(rw http.ResponseWriter, r *http.Request) {
	var request UserAuthRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Printf("DEOCODE ERROR: %s", err)
		rw.WriteHeader(http.StatusBadRequest)
	} else {
		department := handler.Db.GetEmployeeDepartment(r.Context(), request.Login)
		if department.Err != nil {
			log.Printf("GET DEPARTMENT ERROR: %s", err)
			rw.WriteHeader(http.StatusBadRequest)
		} else {
			employees, err := handler.Db.GetDepartmentEmployees(r.Context(), department.Val.(uint64))
			if err != nil {
				log.Printf("GET EMPLOYEES ERROR: %s", err)
				rw.WriteHeader(http.StatusBadRequest)
			} else {
				response := GetDepartmentEmployeesResponse{Employees: employees}
				err = json.NewEncoder(rw).Encode(&response)
				rw.WriteHeader(http.StatusOK)
			}
		}
	}
}
