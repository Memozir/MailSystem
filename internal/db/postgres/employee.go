package db

import (
	"context"
)

func (db *PostgresDB) CreateEmployee(
	ctx context.Context, userId string, roleId string) (employeeId uint8, err error) {
	query := `
			INSERT INTO employee("user", "role") VALUES($1, $2) RETURNING id;
		`
	err = db.connPool.QueryRow(ctx, query, userId, roleId).Scan(employeeId)

	return employeeId, err
}
