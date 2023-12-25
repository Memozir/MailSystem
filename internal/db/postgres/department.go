package db

import (
	"context"
	"github.com/jackc/pgx/v5"
	"mail_system/internal/model"
)

func (db *PostgresDB) GetDepartmentByReceiver(ctx context.Context, receiverId uint64) (uint64, error) {
	query := `
		SELECT a.department
		FROM client as c
		INNER JOIN address as a
		    ON c.address = a.id
		WHERE c.id = $1
	`

	var departmentId uint64
	err := db.connPool.QueryRow(ctx, query, receiverId).Scan(&departmentId)

	return departmentId, err
}

func (db *PostgresDB) AddPackageToStorehouse(
	ctx context.Context,
	departmentId uint64,
	packageId uint64,
	isImport bool) error {

	query := `
		INSERT INTO storehouse(department, package, is_import, is_export)
		VALUES ($1, $2, $3, $4)
	`
	var err error

	if isImport {
		_, err = db.connPool.Exec(ctx, query, departmentId, packageId, true, false)
	} else {
		_, err = db.connPool.Exec(ctx, query, departmentId, packageId, false, true)
	}

	return err
}

func (db *PostgresDB) GetAllDepartments(ctx context.Context) ([]model.Department, error) {
	query := `
		SELECT d.id index, a.name
		FROM department d
		INNER JOIN address a ON  a.department = d.id;
	`

	rows, err := db.connPool.Query(ctx, query)
	departments, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.Department])

	return departments, err
}

func (db *PostgresDB) GetClientDepartments(ctx context.Context, clientId uint64) ([]model.Department, error) {
	query := `
		SELECT d.id index, a.name
		FROM department d 
		INNER JOIN address a on d.id = a.department
		WHERE d.id = (
		    SELECT d.id
		    FROM client c 
		    INNER JOIN address a ON c.address = a.id
		    INNER JOIN department d ON a.department = d.id
		    WHERE c.id = $1
		);
	`

	rows, err := db.connPool.Query(ctx, query, clientId)
	departments, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.Department])

	return departments, err
}
func (db *PostgresDB) GetEmployeeDepartments(ctx context.Context, employeeId uint64) ([]model.Department, error) {
	query := `
		SELECT d.id index, a.name
		FROM department d 
		INNER JOIN address a on d.id = a.department
		INNER JOIN employee e on d.id = e.department
		WHERE d.id = e.department AND e.id = $1
	`

	rows, err := db.connPool.Query(ctx, query, employeeId)
	departments, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.Department])

	return departments, err
}
