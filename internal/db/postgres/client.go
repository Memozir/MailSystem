package db

import (
	"context"
)

func (db *PostgresDB) CreateClient(ctx context.Context, userId uint8, addressName string, apartment string) error {
	addressId, err := db.GetAddressByName(ctx, addressName, apartment)
	query := `INSERT INTO client("user", address) VALUES($1, $2)`

	_, err = db.connPool.Exec(ctx, query, userId, addressId)

	if err != nil {
		return err
	}

	return nil
}

func (db *PostgresDB) AddPackageToClient(ctx context.Context, clientId uint64, packageId uint64) error {
	query := `
		INSERT INTO client_package(client, package) VALUES($1, $2);
	`

	_, err := db.connPool.Query(ctx, query, clientId, packageId)

	return err
}
