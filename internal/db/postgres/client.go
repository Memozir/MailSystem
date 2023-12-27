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

func (db *PostgresDB) AddPackageToClient(ctx context.Context, clientId uint64, packageId uint64) error {
	query := `
		INSERT INTO client_package(client, package) VALUES($1, $2);
	`

	_, err := db.connPool.Query(ctx, query, clientId, packageId)

	return err
}

func (db *PostgresDB) GetClientByLogin(ctx context.Context, login string) (uint64, error) {
	query := `
		SELECT c.id
		FROM client c
		INNER JOIN "user" u
			ON c."user" = u.id
		WHERE u.login = $1;
	`

	var clientId uint64
	err := db.connPool.QueryRow(ctx, query, login).Scan(&clientId)

	return clientId, err
}
