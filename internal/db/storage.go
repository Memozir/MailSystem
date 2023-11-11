package db

type Storage interface {
	Reset()
}

// CreateUser(
// 	first_name string,
// 	second_name string,
// 	phone string,
// 	pass string,
// 	birth string)
// GetUserById(id string) *model.User
