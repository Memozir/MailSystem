package db

import (
	"context"
	"log"
	"strconv"

	"github.com/jackc/pgx/v5"

	"mail_system/internal/model"
)

func (db *PostgresDB) CreateUser(
	ctx context.Context,
	firstName string,
	secondName string,
	login string,
	phone string,
	pass string,
	birth string) uint8 {

	var userId uint8
	err := db.connPool.QueryRow(ctx, `
	INSERT INTO "user"(phone, login, pass, first_name, last_name, birth_date)
	VALUES($1, $2, $3, $4, $5, $6) RETURNING id;
	`, phone, login, pass, firstName, secondName, birth).Scan(&userId)

	if err != nil {
		log.Printf("Falied create user: %s", err.Error())
	}

	return userId
}

func (db *PostgresDB) GetUserById(id string) (*model.User, error) {
	ctx := context.TODO()
	idInt, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		log.Println(err)
	}

	res, err := db.connPool.Query(ctx, `
		SELECT id, phone, pass, first_name, last_name, birth_date::text
		FROM "user"
		WHERE id=$1
	`, idInt)

	if err != nil {
		log.Printf("Get user by id failed: %s", err)
	} else {
		user, err := pgx.CollectOneRow(res, pgx.RowToStructByPos[model.User])

		if err != nil {
			log.Printf("Get user by id query failed: %s", err)
		}

		return &user, err
	}

	return nil, err
}

func (db *PostgresDB) AuthUser(context context.Context, phone string) (bool, error) {
	query := `
		SELECT EXISTS(
			SELECT 1
			FROM client AS c
				INNER JOIN "user" AS u
					ON c.user = u.id
			WHERE u.phone=$1
		);
		`
	var exists bool
	err := db.connPool.QueryRow(context, query, phone).Scan(&exists)

	if err != nil {
		return exists, err
	}

	return exists, nil
}
