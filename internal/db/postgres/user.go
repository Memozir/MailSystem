package db

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/jackc/pgx/v5"

	"mail_system/internal/model"
)

func (db *PostgresDB) CreateUser(
	ctx context.Context,
	cancelFunc context.CancelFunc,
	firstName string,
	secondName string,
	middleName string,
	login string,
	pass string,
	birth string) ResultDB {

	var userId uint8
	err := db.connPool.QueryRow(ctx, `
	INSERT INTO "user"(login, pass, first_name, last_name, middle_name, birth_date)
	VALUES($1, $2, $3, $4, $5, $6) RETURNING id;
	`, login, pass, firstName, secondName, middleName, birth).Scan(&userId)

	if err != nil {
		log.Printf("Falied create user: %s", err.Error())
		cancelFunc()
	}
	fmt.Println(userId)
	return ResultDB{userId, err}
}

func (db *PostgresDB) GetUserById(id string) ResultDB {
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

		return ResultDB{user, err}
	}

	return ResultDB{Err: err}
}

func (db *PostgresDB) AuthUser(context context.Context, login string, pass string) (bool, error) {
	query := `
		SELECT EXISTS(
			SELECT 1
			FROM client AS c
				INNER JOIN "user" AS u
					ON c.user = u.id
			WHERE u.login=$1 and u.pass=$2
		);
		`
	var exists bool
	err := db.connPool.QueryRow(context, query, login, pass).Scan(&exists)

	if err != nil {
		return exists, err
	}

	return exists, nil
}
