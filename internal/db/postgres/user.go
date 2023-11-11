package db

import (
	"context"
	"log"
	"mail_system/internal/model"
	"strconv"

	"github.com/jackc/pgx/v5"
)

type User struct {
	Db *PostgresDB
}

func (user *User) CreateUser(
	first_name string,
	second_name string,
	phone string,
	pass string,
	birth string) {
	ctx := context.TODO()
	pgcon, err := user.Db.connPool.Exec(ctx, `
		INSERT INTO Users(phone, pass, first_name, second_name, birth_date)
		VALUES($1, $2, $3, $4, $5)
		`, phone, pass, first_name, second_name, birth)

	if err != nil {
		log.Printf("Falied create user: %s", err.Error())
	}
	log.Printf("PGCON: %s", pgcon)
}

func (user *User) GetUserById(id string) *model.User {
	ctx := context.TODO()
	idInt, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		log.Println(err)
	}

	res, err := user.Db.connPool.Query(ctx, `
		SELECT id, phone, pass, first_name, second_name, birth_date::text
		FROM Users
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
