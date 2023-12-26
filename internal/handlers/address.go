package handlers

import (
	"encoding/json"
	"log"
	"mail_system/internal/config"
	"mail_system/internal/model"
	"net/http"
)

func (handler *MailHandlers) CreateAddressHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("Address registration handler")

	var request ManageAddressJSON
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		log.Printf("ADDRES DECODE ERROR: %s", err)
		rw.WriteHeader(http.StatusBadRequest)
	} else {
		employee := handler.Db.GetEmployeeByLogin(r.Context(), request.User.Login)
		if employee.Val.(model.Employee).RoleCode >= int8(config.AdminRole) {
			err = handler.Db.CreateAddress(r.Context(),
				employee.Val.(model.Employee).DepartmentId, request.AddressName)
			if err != nil {
				log.Printf("ADDRES CREATION ERROR: %s", err)
				rw.WriteHeader(http.StatusBadRequest)
			} else {
				log.Printf("ADDRESS <%s> WAS CREATED", request.AddressName)
				rw.WriteHeader(http.StatusOK)
			}
		} else {
			log.Println("NOT ENOUGH RIGHTS")
			rw.WriteHeader(http.StatusBadRequest)
		}
	}

	// log.Println(request)
	// contextCreateAddress, cancelCreateAddress := context.WithTimeout(r.Context(), time.Second*2)
	// defer cancelCreateAddress()

	// err = handler.Db.CreateAddress(contextCreateAddress, request.Name)

	// if err != nil {
	// 	log.Printf("Address was not created: %s", err.Error())
	// }

	// rw.Header().Set("Content-type", "application/json")
}
