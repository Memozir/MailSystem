package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"mail_system/internal/config"

	pgx "github.com/jackc/pgx/v5/pgxpool"
)

type PostgresDB struct {
	connPool *pgx.Pool
	cfg      *config.ConfigDb
}

func NewDb(ctx context.Context) (db *PostgresDB) {
	db = new(PostgresDB)
	var err error

	db.cfg = &config.ConfigDb{
		Host:         os.Getenv("DB_HOST"),
		Port:         os.Getenv("DB_PORT"),
		Pass:         os.Getenv("DB_PASSWORD"),
		User:         os.Getenv("DB_USER"),
		DbName:       os.Getenv("DB_NAME"),
		SSLMode:      os.Getenv("DB_SSL_MODE"),
		MaxPoolConns: os.Getenv("DB_MAX_CONN_POOLS"),
	}

	url := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s pool_max_conns=%s",
		db.cfg.User,
		db.cfg.Pass,
		db.cfg.Host,
		db.cfg.Port,
		db.cfg.DbName,
		db.cfg.SSLMode,
		db.cfg.MaxPoolConns,
	)

	db.connPool, err = pgx.New(ctx, url)

	if err != nil {
		log.Fatalf("Failed connection to Postgres: %s", err)
		return db
	}

	log.Printf("Connection to database on %s was Success", db.cfg.Host)
	return db
}

func (db *PostgresDB) CreateUser(
	first_name string,
	second_name string,
	phone string,
	pass string,
	birth string,
) {
	ctx := context.TODO()
	pgcon, err := db.connPool.Exec(ctx, `
		INSERT INTO Users(phone, pass, first_name, second_name, birth_date)
		VALUES($1, $2, $3, $4, $5)
		`, phone, pass, first_name, second_name, birth)

	if err != nil {
		log.Printf("Falied create user: %s", err.Error())
	}
	log.Printf("PGCON: %s", pgcon)
}
