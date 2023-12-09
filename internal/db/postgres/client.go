package db

import (
	"context"
)

func (db *PostgresDB) CreateClient(ctx context.Context, userId uint8, addressName string) (uint8, error) {
	addressId, err := db.GetAddressByName(ctx, addressName)

	if err != nil {
		return addressId, err
	}

	return addressId, nil
}
