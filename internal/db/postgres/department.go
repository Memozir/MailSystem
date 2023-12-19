package db

import "context"

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
		_, err = db.connPool.Query(ctx, query, departmentId, packageId, true, false)
		if err != nil {
			return err
		}
	} else {
		_, err = db.connPool.Query(ctx, query, departmentId, packageId, false, true)
		if err != nil {
			return err
		}
	}

	return err
}
