package db

import (
	"context"
	"mail_system/internal/config"
	"mail_system/internal/model"
	"strings"

	"github.com/jackc/pgx/v5"
)

type Tariff struct {
	Id    uint64 `db:"id"`
	Price uint8  `db:"price"`
}

func (db *PostgresDB) ProducePaymentInfo(ctx context.Context, packageId uint64, packageType int, weight int) error {
	query := `
		SELECT price, id FROM tarrif WHERE type = $1 and weight = $2;
	`

	var tariff Tariff
	row, err := db.connPool.Query(ctx, query, packageType, weight)
	if err != nil {
		return err
	}
	tariff, err = pgx.CollectOneRow(row, pgx.RowToStructByName[Tariff])
	if err != nil {
		return err
	}

	err = db.SetPackagePaymentInfo(ctx, packageId, tariff.Id)
	if err != nil {
		return err
	}

	return nil
}

func (db *PostgresDB) SetPackagePaymentInfo(ctx context.Context, packageId uint64, tariffId uint64) error {
	query := `
		INSERT INTO payment_info(package, tarrif) VALUES($1, $2);
	`

	_, err := db.connPool.Query(ctx, query, packageId, tariffId)

	return err
}

func (db *PostgresDB) AddEmployeeToPackageResponsibleList(ctx context.Context, employeeId uint64, packageId uint64) error {
	query := `
		INSERT INTO public.employee_package(employee, package) VALUES($1, $2);
	`

	_, err := db.connPool.Query(ctx, query, employeeId, packageId)

	return err
}

func (db *PostgresDB) CreatePackage(
	ctx context.Context,
	weight int,
	packageType int,
	senderId uint64,
	receiverId uint64,
	departmentReceiver uint64,
	createDate string,
	deliverDate string) (uint64, error) {

	query := `
		INSERT INTO package(
		                    weight,
		                    sender,
		                    receiver,
		                    department_receiver,
		                    "type",
		                    create_date,
		                    deliver_date)
		VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id
	`

	var packageId uint64
	err := db.connPool.QueryRow(
		ctx,
		query,
		weight,
		senderId,
		receiverId,
		departmentReceiver,
		packageType,
		createDate,
		deliverDate,
	).Scan(&packageId)

	return packageId, err
}

/*
	Id             uint64 `db:"id"`
	Status         int    `db:"status"`
	Weight         int    `db:"weight"`
	Sender         uint64 `db:"sender"`
	Receiver       uint64 `db:"receiver"`
	Courier        uint64 `db:"courier"`
	DateOfCreation string `db:"create_date"`
	DateOfReceipt  string `db:"deliver_date"`
	NumDepartment  uint64 `db:"department_receiver"`
	Address        uint64 `db:"address"
*/

func (db *PostgresDB) GetEmployeePackages(
	ctx context.Context,
	employeeId uint64,
	employeeRole uint8) ([]model.Package, error) {
	query := `
		SELECT 
		    DISTINCT p.id,
		    p.status,
		    p.weight,
		    u_rec.login as receiver,
		    u_send.login as sender,
		    CAST(p.create_date AS TEXT),
		 	CAST(p.deliver_date AS TEXT),
		    p.department_receiver,
		    CONCAT_WS(' ', addr.name, send.apartment) sender_address,
		    COALESCE(
		    	(
		    	SELECT u_cur.login
		    	FROM employee_package as emp
		    	INNER JOIN employee as cur_emp ON emp.employee = cur_emp.id
				INNER JOIN "user" u_cur ON cur_emp."user" = u_cur.id
		    	WHERE cur_emp.role = 1 and emp.package = p.id),
		    	'') as courier,
		    t.type
		FROM employee_package as ep
			INNER JOIN package as p ON ep.package = p.id
			INNER JOIN client rec ON p.receiver = rec.id
			INNER JOIN client send ON p.sender = send.id
			INNER JOIN address addr ON addr.id = rec.address
			INNER JOIN "user" u_rec ON rec."user" = u_rec.id
			INNER JOIN "user" u_send ON send."user" = u_send.id
			INNER JOIN payment_info pi ON pi.package = p.id
			INNER JOIN tarrif as t ON pi.tarrif = t.id
		WHERE ep.employee = $1
	`
	var queryBuilder strings.Builder
	queryBuilder.WriteString(query)
	var rows pgx.Rows
	var err error

	if employeeRole == config.CourierRole {
		queryBuilder.WriteString(` and p.status = $2`)
		rows, err = db.connPool.Query(ctx, queryBuilder.String(), employeeId, config.PACKAGE_STATUS_DELIVERY)
	} else {
		rows, err = db.connPool.Query(ctx, query, employeeId)
	}

	packages, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.Package])

	return packages, err
}

func (db *PostgresDB) GetCourierDeliverPackages(ctx context.Context, departmentId uint64) ([]model.Package, error) {
	query := `
		SELECT
			DISTINCT p.id,
		    p.status,
		    p.weight,
		    u2_rec.login as receiver,
		    u_send.login as sender,
		    CAST(p.create_date AS TEXT),
		 	CAST(p.deliver_date AS TEXT),
		    p.department_receiver,
		    CONCAT_WS(' ', addr.name, c2.apartment) sender_address,
		    COALESCE(
		    	(
		    	SELECT id
		    	FROM employee_package as emp
		    	INNER JOIN employee as cur_emp ON emp.employee = cur_emp.id
		    	WHERE cur_emp.role = 1 and emp.package = p.id),
		    	0) as courier,
		    p.type
		FROM storehouse s 
			INNER JOIN package p ON s.package = p.id
			INNER JOIN client c on c.id = p.receiver
			INNER JOIN client c2 on c2.id = p.sender
			INNER JOIN "user" u_send on u_send.id = c."user"
			INNER JOIN "user" u2_rec on u2_rec.id = c2."user"
			INNER JOIN public.address addr on addr.id = c.address
		WHERE s.department = $1 and p.status = $2
	`

	rows, err := db.connPool.Query(ctx, query, departmentId, config.PACKAGE_STARUS_DELIVERY_AWAITED)

	packages, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.Package])

	return packages, err
}

func (db *PostgresDB) ChangePackageStatus(ctx context.Context, packageID uint64, status uint8) error {
	query := `
		UPDATE package
		SET status = $1
		WHERE id = $2
	`

	_, err := db.connPool.Exec(ctx, query, status, packageID)

	return err
}

func (db *PostgresDB) GetClientPackages(ctx context.Context, clientId uint64) ([]model.Package, error) {
	query := `
		SELECT
		    p.id,
		    p.status,
		    p.weight,
		    u2_rec.login as receiver,
		    u_send.login as sender,
		    CAST(p.create_date AS TEXT),
		 	CAST(p.deliver_date AS TEXT),
		    p.department_receiver,
		    CONCAT_WS(' ', addr.name, c2.apartment) sender_address,
		    COALESCE(
		    	(
		    	SELECT id
		    	FROM employee_package as emp
		    	INNER JOIN employee as cur_emp ON emp.employee = cur_emp.id
		    	WHERE cur_emp.role = 1 and emp.package = p.id),
		    	0) as courier,
		    p.type
		FROM package p
			INNER JOIN client c on c.id = p.receiver
			INNER JOIN client c2 on c2.id = p.sender
			INNER JOIN "user" u_send on u_send.id = c2."user"
			INNER JOIN "user" u2_rec on u2_rec.id = c."user"
			INNER JOIN public.address addr on addr.id = c.address
		WHERE p.sender = $1 and p.status < $2
	`

	rows, err := db.connPool.Query(ctx, query, clientId, config.PACKAGE_STATUS_RECEIVED)

	packages, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.Package])

	return packages, err
}
