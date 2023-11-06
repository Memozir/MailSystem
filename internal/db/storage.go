package db

import "mail_system/internal/model"

type Storage interface {
	CreateUser(
		first_name string,
		second_name string,
		phone string,
		pass string,
		birth string)
	GetUserById(id string) *model.User
	Reset()
}
