package db

type Storage interface {
	createTables() error
}
