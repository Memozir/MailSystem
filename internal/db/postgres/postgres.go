package db

import (
	"context"
	"fmt"
	"log"
	"os"

	pgxPool "github.com/jackc/pgx/v5/pgxpool"

	"mail_system/internal/config"
)

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

	// Host:         utils.GetEnvOrDefault("DB_HOST", "localhost"),
	// Port:         utils.GetEnvOrDefault("DB_PORT", "5431"),
	// Pass:         utils.GetEnvOrDefault("DB_PASSWORD", "postgres"),
	// User:         utils.GetEnvOrDefault("DB_USER", "postgres"),
	// DbName:       utils.GetEnvOrDefault("DB_NAME", "mail_system_db"),
	// SSLMode:      utils.GetEnvOrDefault("DB_SSL_MODE", "disable"),
	// MaxPoolConns: utils.GetEnvOrDefault("DB_MAX_CONN_POOLS", "10")}

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
		log.Fatalf("Failed connection to Postgres: %s", err)
		return db
	}

	log.Printf("Connection to database on %s was Success", db.cfg.Host)
	return db
}
