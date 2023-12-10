package db

import (
	"context"
)

func (db *PostgresDB) CreateClient(ctx context.Context, userId uint8, addressName string) error {
	addressId, err := db.GetAddressByName(ctx, addressName)
	query := `INSERT INTO client("user", address) VALUES($1, $2)`

	_, err = db.connPool.Exec(ctx, query, userId, addressId)

	if err != nil {
		return err
	}

	return nil
}
