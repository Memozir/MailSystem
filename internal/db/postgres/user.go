package db

import (
	"context"
	"github.com/jackc/pgx/v5"
	"mail_system/internal/model"
)

func (db *PostgresDB) CreateUser(ctx context.Context,
	firstName string,
	secondName string,
	middleName string,
	login string,
	pass string,
	birth string) (ResultDB, error) {

	var userId uint64
	err := db.connPool.QueryRow(ctx, `
	INSERT INTO "user"(login, pass, first_name, last_name, middle_name, birth_date)
	VALUES($1, $2, $3, $4, $5, $6) RETURNING id;
	`, login, pass, firstName, secondName, middleName, birth).Scan(&userId)

	return ResultDB{userId, err}, err
}

func (db *PostgresDB) GetUserIdByLogin(ctx context.Context, login string) (ResultDB, error) {
	query := `
		SELECT id
		FROM client
		WHERE login=$1`

	var userId uint64
	err := db.connPool.QueryRow(ctx, query, login).Scan(&userId)

	if err != nil {
		return ResultDB{}, err
	}

	return ResultDB{userId, err}, err
}

type SenderReceiverRes struct {
	Sender   uint64 `db:"sender"`
	Receiver uint64 `db:"receiver"`
}

func (db *PostgresDB) GetSenderReceiverIdByLogin(ctx context.Context, senderLogin string, receiverLogin string) (ResultDB, error) {
	query := `
		SELECT
		(
			SELECT c1.id
			FROM client as c1
			INNER JOIN "user" u1
				ON c1."user" = u1.id
			WHERE u1.login = $1
		) sender,
		(
			SELECT c2.id
			FROM client as c2
			INNER JOIN "user" u2
				ON c2."user" = u2.id
			WHERE u2.login = $2
		) receiver;
		`

	var senderReceiver SenderReceiverRes
	row, err := db.connPool.Query(ctx, query, senderLogin, receiverLogin)
	if err != nil {
		return ResultDB{}, err
	}
	senderReceiver, err = pgx.CollectOneRow(row, pgx.RowToStructByName[SenderReceiverRes])
	if err != nil {
		return ResultDB{}, err
	}

	return ResultDB{senderReceiver, err}, err
}

func (db *PostgresDB) AuthUser(ctx context.Context, login string, pass string) (ResultDB, error) {
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
	err := db.connPool.QueryRow(ctx, query, login, pass).Scan(&userAuth.ClientId, &userAuth.RoleCode)

	return ResultDB{Val: userAuth, Err: err}, err
}
