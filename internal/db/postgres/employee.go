package db

import (
	"context"
)

func (db *PostgresDB) CreateEmployee(
	ctx context.Context, userId uint8, roleId uint8) ResultDB {
	query := `
			INSERT INTO employee("user", "role") VALUES($1, $2) RETURNING id;
		`
	var employeeId uint8
	err := db.connPool.QueryRow(ctx, query, userId, roleId).Scan(&employeeId)

	return ResultDB{Err: err, Val: employeeId}
}
