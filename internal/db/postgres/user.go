package db

import (
	"context"
	"log"
	"strconv"

	"github.com/jackc/pgx/v5"

	"mail_system/internal/model"
)

func (db *PostgresDB) CreateUser(
	first_name string,
	second_name string,
	phone string,
	pass string,
	birth string) uint8 {
	ctx := context.TODO()

	// pgcon, err := db.connPool.Exec(ctx, `
	// 	INSERT INTO "user"(phone, pass, first_name, last_name, birth_date)
	// 	VALUES($1, $2, $3, $4, $5)
	// 	`, phone, pass, first_name, second_name, birth)

	var userId uint8
	err := db.connPool.QueryRow(ctx, `
	INSERT INTO "user"(phone, pass, first_name, last_name, birth_date)
	VALUES($1, $2, $3, $4, $5) RETURNING id;
	`, phone, pass, first_name, second_name, birth).Scan(&userId)

	if err != nil {
		log.Printf("Falied create user: %s", err.Error())
	}
	// log.Printf("PGCON: %s", pgcon)

	return userId
}

func (db *PostgresDB) GetUserById(id string) *model.User {
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

		return &user
	}

	return nil
}
