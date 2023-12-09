package db

import (
	"context"
	"log"
)

func (db *PostgresDB) CreateAddress(ctx context.Context, name string) error {
	_, err := db.connPool.Query(ctx, `INSERT INTO address("name") VALUES($1);`, name)

	if err != nil {
		log.Printf("Address was not created: %s", err.Error())
		return err
	}

	return nil
}

func (db *PostgresDB) GetAddressByName(ctx context.Context, name string) (uint8, error) {
	var addressId uint8
	query := `SELECT id FROM "address" WHERE "name" = $1;`
	err := db.connPool.QueryRow(ctx, query, name).Scan(&addressId)

	if err != nil {
		return addressId, err
	}

	return addressId, nil
}
