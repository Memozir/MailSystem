package db

import (
	// "context"
	"context"
	"fmt"
	"log"

	// "net/url"

	pgx "github.com/jackc/pgx/v5"
)

type PostgresDB struct {
	db *pgx.Conn
}

func NewPostgresDb() (*PostgresDB, error) {
	// "postgres://username:password@localhost:5432/database_name"
	// config := pgx.ConnConfig{
	// 	Host:      "localhost",
	// 	Port:      5432,
	// 	Database:  "url_shorter_db",
	// 	User:      "postgres",
	// 	Password:  "postgres",
	// 	TLSConfig: nil,
	// }

	urlConnect := "postgres://postgres:postgres@localhost:5431/url_shorter_db"
	conn, err := pgx.Connect(context.Background(), urlConnect)

	if err != nil {
		log.Fatalf("Connection to postgres db %s failed", urlConnect)
	}

	return &PostgresDB{db: conn}, nil
}

func (db *PostgresDB) createTables() error {
	fmt.Printf("Tables was created")

	return nil
}
