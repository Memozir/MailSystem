package db

import (
	"context"
	"fmt"
	"log"
	"mail_system/internal/model"
	"os"

	pgxPool "github.com/jackc/pgx/v5/pgxpool"

	"mail_system/internal/config"
)

type Storage interface {
	Reset()
	CreateUser(ctx context.Context,
		firstName string,
		secondName string,
		login string,
		phone string,
		pass string,
		birth string) uint8
	CreateEmployee(ctx context.Context, userId uint8, roleId uint8) (employeeId uint8, err error)
	CreateRole(ctx context.Context, code uint8, name string) error
	CreateAddress(ctx context.Context, name string) error
	GetAddressByName(ctx context.Context, name string) (uint8, error)
	CreateClient(ctx context.Context, userId uint8, addressName string) error
	GetRoleByName(ctx context.Context, roleName string) (uint8, error)
	GetUserById(id string) (*model.User, error)
	AuthUser(context context.Context, phone string) (bool, error)
}

type PostgresDB struct {
	connPool *pgxPool.Pool
	cfg      *config.ConfigDb
}

func (db *PostgresDB) Reset() {
	db.connPool.Reset()
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
		MaxPoolConns: os.Getenv("DB_MAX_CONN_POOLS")}

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

	db.connPool, err = pgxPool.New(ctx, url)

	if err != nil {
		// log.Fatalf("Failed connection to Postgres: %s", err)
		log.Panicf("Failed connection to Postgres: %s", err)
	}
	if err = db.connPool.Ping(ctx); err != nil {
		log.Panicf("Failed connection to Postgres: %s", err)
	}

	log.Printf("Connection to database on %s was Success", db.cfg.Host)
	return db
}
