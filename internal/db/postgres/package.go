package db

import (
	"context"
)

type Tariff struct {
	Id    uint64 `db:"id"`
	Price uint8  `db:"id"`
}

func (db *PostgresDB) ProducePaymentInfo(ctx context.Context, packageId uint64, packageType int, weight int) error {
	query := `
		SELECT price, id FROM tarrif WHERE type = $1 and weight = $2;
	`

	var tariff Tariff
	err := db.connPool.QueryRow(ctx, query, packageType, weight).Scan(&tariff)

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
		VALUES ($1, $2, $3, $4, $5, $6, $7)
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