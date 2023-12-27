package db

import (
	"context"
	"mail_system/internal/config"
	"mail_system/internal/model"

	"github.com/jackc/pgx/v5"
)

func (db *PostgresDB) CreateEmployee(
	ctx context.Context, userId uint64, departmentId uint64, roleCode uint8) (ResultDB, error) {
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

	return ResultDB{Err: err, Val: employeeId}, err
}

func (db *PostgresDB) DeleteEmployee(ctx context.Context, employeeId uint64) error {
	query := `
		DELETE FROM employee WHERE id = $1;
	`
	_, err := db.connPool.Exec(ctx, query, employeeId)
	return err
}

func (db *PostgresDB) GetEmployeeByLogin(ctx context.Context, login string) (ResultDB, error) {
	query := `
		SELECT e.id, e."user", e."role", e.department
		FROM employee e
		INNER JOIN "user" u
			ON e."user" = u.id
		WHERE u.login = $1;
	`

	var employee model.Employee
	row, err := db.connPool.Query(ctx, query, login)
	if err != nil {
		return ResultDB{Val: employee, Err: err}, err
	}
	employee, err = pgx.CollectOneRow(row, pgx.RowToStructByName[model.Employee])
	if err != nil {
		return ResultDB{Val: employee, Err: err}, err
	}

	return ResultDB{Val: employee, Err: err}, err
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

func (db *PostgresDB) GetEmployeeDepartmentByRole(ctx context.Context, departmentId uint64, role int) (uint64, error) {
	query := `
		SELECT e.id
		FROM employee e
		INNER JOIN role r ON e.role = r.code
		WHERE e.department = $1 and r.code = $2 
	`

	var employeeId uint64
	err := db.connPool.QueryRow(ctx, query, departmentId, role).Scan(&employeeId)

	return employeeId, err
}

func (db *PostgresDB) CheckAdminAddress(ctx context.Context, adminId uint64, addressName string) (bool, error) {
	query := `
		SELECT EXISTS(
			SELECT a.id
			FROM address a
			WHERE a.department = (
				SELECT d.id
				FROM employee e
				INNER JOIN department d ON e.department = d.id
				WHERE e.id = $1 and a."name" = $2
			)
		)
	`
	var res bool
	err := db.connPool.QueryRow(ctx, query, adminId, addressName).Scan(&res)

	return res, err
}
