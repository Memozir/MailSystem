package db

import (
	"context"
)

func (db *PostgresDB) CreateClient(ctx context.Context, userId uint8, addressName string, apartment string) error {
	addressId, err := db.GetAddressByName(ctx, addressName)
	query := `INSERT INTO client("user", address, apartment) VALUES($1, $2, $3)`

	_, err = db.connPool.Exec(ctx, query, userId, addressId, apartment)

	if err != nil {
		return err
	}

	return nil
}
