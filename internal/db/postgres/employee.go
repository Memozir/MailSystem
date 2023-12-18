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

func (db *PostgresDB) GetEmployeeByLogin(ctx context.Context, login string) (uint64, error) {
	query := `
		SELECT e.id
		FROM employee e
		INNER JOIN "user" u
			ON e."user" = u.id
		WHERE u.login = $1;
	`

	var employeeId uint64
	err := db.connPool.QueryRow(ctx, query, login).Scan(&employeeId)

	return employeeId, err
}
