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

func (db *PostgresDB) AuthUser(ctx context.Context, login string, pass string) ResultDB {
	query := `
		SELECT
		    COALESCE(c.id, 0) as client_id,
		    COALESCE(emp.code, 0) as role_code
		FROM "user" AS u
         LEFT JOIN client AS c
                    ON u.id = c.user
         LEFT JOIN (SELECT "user" user_id, r.code code
                  FROM employee as e
                    INNER JOIN "role" as r
                      on e.role = r.code) as emp
             ON u.id = emp.user_id
		WHERE u.login=$1 and u.pass=$2;
		`
	var userAuth model.UserAuth
	err := db.connPool.QueryRow(ctx, query, login, pass).Scan(&userAuth.RoleCode, &userAuth.ClientId)

	return ResultDB{Val: userAuth, Err: err}
}
