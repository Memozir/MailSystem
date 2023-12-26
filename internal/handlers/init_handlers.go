package handlers

import (
	"context"

	db "mail_system/internal/db/postgres"

	"github.com/gorilla/mux"

	httpSwagger "github.com/swaggo/http-swagger"
	//_ "mail_system/internal/handlers/docs"
	_ "mail_system/docs"
)

type MailHandlers struct {
	Context context.Context
	Db      db.Storage
}

func NewMailHandler(db db.Storage) *MailHandlers {
	return &MailHandlers{
		Db: db,
	}
}

// LoadHandlers @title MailSystem API
// @version 1.0
// @description This is a service for managing mail system
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func (handler *MailHandlers) LoadHandlers() *mux.Router {
	router := mux.NewRouter()

	// Adding handlers
	router.HandleFunc("/register/client", handler.RegisterClientHandler).Methods("POST")
	router.HandleFunc("/register/employee", handler.RegisterEmployeeHandler).Methods("POST")
	//router.HandleFunc("/user/{id}", handler.GetUserHandler).Methods("GET")
	router.HandleFunc("/create/address", handler.CreateAddressHandler).Methods("POST")
	router.HandleFunc("/auth/user", handler.AuthUserHandler).Methods("POST")
	router.HandleFunc("/create/role", handler.CreateRoleHandler).Methods("POST")
	router.HandleFunc("/create/package", handler.CreateDepartmentPackageHandler).Methods("POST")
	router.HandleFunc("/get/packages", handler.GetEmployeePackages).Methods("GET")
	router.HandleFunc("/get/departments", handler.GetAllDepartments).Methods("GET")
	router.HandleFunc("/get/client/departments", handler.GetClientDepartments).Methods("GET")
	router.HandleFunc("/get/employee/departments", handler.GetEmployeeDepartments).Methods("GET")
	router.HandleFunc("/delete/address", handler.DeleteAddressByAdmin).Methods("DELETE")

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	return router
}
