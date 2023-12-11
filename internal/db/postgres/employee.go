package db

import (
	"context"
)

func (db *PostgresDB) CreateEmployee(
	ctx context.Context, userId uint8, roleId uint8) (employeeId uint8, err error) {
	query := `
			INSERT INTO employee("user", "role") VALUES($1, $2) RETURNING id;
		`
	err = db.connPool.QueryRow(ctx, query, userId, roleId).Scan(employeeId)

	return employeeId, err
}
