package db

import (
	"context"
	"log"
)

func (db *PostgresDB) CreateAddress(ctx context.Context, departmentId uint64, addressName string) error {
	_, err := db.connPool.Query(ctx, `
	INSERT INTO address("name", department) VALUES($1, $2);`, addressName, departmentId)

	if err != nil {
		log.Printf("Address was not created: %s", err.Error())
		return err
	}

	return nil
}

func (db *PostgresDB) GetAddressByName(ctx context.Context, addressName string) (uint8, error) {
	var addressId uint8
	query := `SELECT id FROM "address" WHERE "name" = $1;`
	err := db.connPool.QueryRow(ctx, query, addressName).Scan(&addressId)

	if err != nil {
		return addressId, err
	}

	return addressId, nil
}

func (db *PostgresDB) DeleteAddress(ctx context.Context, addressName string) error {
	query := `
		DELETE FROM address WHERE "name" = $1;
	`

	_, err := db.connPool.Exec(ctx, query, addressName)

	return err
}
