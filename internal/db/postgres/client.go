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
