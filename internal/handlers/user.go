package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mail_system/internal/config"
	"mail_system/internal/model"
	utils "mail_system/internal/utils/auth"
	"net/http"
)

type UserJSON struct {
	Id         uint64 `json:"id"`
	Phone      string `json:"phone"`
	Login      string `json:"login"`
	Pass       string `json:"pass"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	MiddleName string `json:"middle_name"`
	BirthDate  string `json:"birth_date"`
	Apartment  string `json:"appartment"`
}

func (handler *MailHandlers) RegisterUserHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("User registration handler")

	var userJSON UserJSON
	err := json.NewDecoder(r.Body).Decode(&userJSON)

	if err != nil {
		log.Printf("User decode error: %s", err)
	}

	log.Println(userJSON)
	//contextCreateUser, cancelCreateUser := context.WithTimeout(r.Context(), time.Second*2)
	//defer cancelCreateUser()

	userId, err := handler.Db.CreateUser(
		r.Context(),
		userJSON.FirstName,
		userJSON.SecondName,
		userJSON.MiddleName,
		userJSON.Login,
		userJSON.Pass,
		userJSON.BirthDate)
	rw.Header().Set("Content-type", "application/json")

	fmt.Printf("User id: %d", userId.Val.(uint8))
}

/*
func (handler *MailHandlers) GetUserHandler(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := handler.Db.GetUserById(vars["id"])

	if user.Err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println(user.Val.(model.User))
	rw.WriteHeader(http.StatusOK)
}
*/

type UserAuthRequest struct {
	Login string `json:"login"`
	Pass  string `json:"pass"`
}

type UserAuthResponse struct {
	Role        int8   `json:"role"`
	AccessToken string `json:"token"`
}

// AuthUserHandler AuthUser godoc
// @Summary Create a new UserAuthRequest
// @Description Create a new order with the input payload
// @Tags auth
// @Accept  json
// @Param data body UserAuthRequest true
// @Produce  json
// @Success 200 {object} UserAuthResponse
// @Router /auth/user [post]
func (handler *MailHandlers) AuthUserHandler(rw http.ResponseWriter, r *http.Request) {
	var userAuth UserAuthRequest
	err := json.NewDecoder(r.Body).Decode(&userAuth)

	if err != nil {
		log.Println(err.Error())
		rw.WriteHeader(http.StatusBadRequest)
	}

	//contextAuth, cancelAuth := context.WithTimeout(context.Background(), time.Second*2)
	//defer cancelAuth()
	res, err := handler.Db.AuthUser(context.Background(), userAuth.Login, userAuth.Pass)

	if err != nil {
		log.Println("USER WAS NOT FOUND AND AUTHORIZED")
		log.Println(err.Error())
		rw.WriteHeader(http.StatusBadRequest)
	} else if res.Val.(model.UserAuth).ClientId > 0 {
		log.Println("SUCCESS CLIENT AUTH")

		accessToken, err := utils.GetAccessToken("client")

		if err != nil {
			http.Error(rw, "Token error", http.StatusBadGateway)
			return
		}

		response := UserAuthResponse{Role: config.UserRole, AccessToken: accessToken}
		err = json.NewEncoder(rw).Encode(&response)

		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			log.Println(err.Error())
		}

		rw.WriteHeader(http.StatusOK)
	} else if res.Val.(model.UserAuth).RoleCode != 0 {
		log.Println("SUCCESS EMPLOYEE AUTH")

		accessToken, err := utils.GetAccessToken("employee")

		if err != nil {
			http.Error(rw, "Token error", http.StatusBadGateway)
			return
		}

		response := UserAuthResponse{Role: res.Val.(model.UserAuth).RoleCode, AccessToken: accessToken}
		err = json.NewEncoder(rw).Encode(response)

		if err != nil {
			log.Println(err.Error())
			rw.WriteHeader(http.StatusBadRequest)
		}
	} else {
		log.Println("ERROR AUTH")
		rw.WriteHeader(http.StatusBadRequest)
	}
}

/*
func (handler *MailHandlers) Test1(rw http.ResponseWriter, r *http.Request) {
	_, err := handler.Db.GetSenderReceiverIdByLogin(r.Context(), "qw3rt", "1")

	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println("SUCCESS")
	}
}
*/
