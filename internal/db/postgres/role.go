package db

import (
	"context"
)

func (db *PostgresDB) CreateRole(ctx context.Context, code uint8, name string) ResultDB {
	query := `
		INSERT INTO role VALUES ($1, $2);
	`
	_, err := db.connPool.Exec(ctx, query, code, name)

	return ResultDB{Err: err}
}

func (db *PostgresDB) GetRoleByName(ctx context.Context, roleName string) ResultDB {
	query := `SELECT code FROM "role" WHERE name=$1`
	var roleCode uint8
	err := db.connPool.QueryRow(ctx, query, roleName).Scan(&roleCode)
	return ResultDB{roleCode, err}
}
