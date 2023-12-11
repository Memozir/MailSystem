package db

import (
	"context"
)

func (db *PostgresDB) CreateRole(ctx context.Context, code uint8, name string) error {
	query := `
		INSERT INTO role VALUES ($1, $2);
	`
	_, err := db.connPool.Exec(ctx, query, code, name)

	return err
}
