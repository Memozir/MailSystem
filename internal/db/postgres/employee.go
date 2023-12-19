package db

import (
	"context"
	"mail_system/internal/config"
)

func (db *PostgresDB) CreateEmployee(
	ctx context.Context, userId uint8, departmentId uint64, roleCode uint8) ResultDB {
	query := `
			INSERT INTO employee("user", "role", department) VALUES($1, $2, $3) RETURNING id;
		`
	var employeeId uint8
	err := db.connPool.QueryRow(ctx, query, userId, roleCode, departmentId).Scan(&employeeId)

	if err != nil && roleCode == config.CourierRole {
		query = `
			INSERT INTO delivery_schedule(courier) VALUES ($1)
		`
		_, err = db.connPool.Exec(ctx, query, employeeId)
	}

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

func (db *PostgresDB) GetEmployeeDepartment(ctx context.Context, login string) ResultDB {
	query := `
		SELECT e.department
		FROM employee e
		INNER JOIN "user" u on u.id = e."user"
		WHERE u.login = $1 
	`

	var departmentId uint64
	err := db.connPool.QueryRow(ctx, query, login).Scan(&departmentId)

	return ResultDB{Val: departmentId, Err: err}
}
