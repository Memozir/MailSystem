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

func (db *PostgresDB) GetRoleByName(ctx context.Context, roleName string) (uint8, error) {
	query := `SELECT code FROM "role" WHERE name=$1`
	var roleCode uint8
	err := db.connPool.QueryRow(ctx, query, roleName).Scan(&roleCode)
	return roleCode, err
}
