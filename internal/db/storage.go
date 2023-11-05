package db

type Storage interface {
	CreateUser(
		first_name string,
		second_name string,
		phone string,
		pass string,
		birth string)
}
