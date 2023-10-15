package db

import (
	// "context"
	"context"
	"fmt"
	"log"

	// "net/url"

	// _ "github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

type PostgresDB struct {
	// conn *pgx.Conn
	conn *sqlx.DB
}

func NewPostgresDb() (*PostgresDB, error) {
	// urlConnect := "postgres://postgres:postgres@localhost:5431/url_shorter_db"
	// conn, err := pgx.Connect(context.Background(), urlConnect)

	// if err != nil {
	// 	log.Fatalf("Connection to postgres db %s failed", urlConnect)
	// }

	// conn, err := sqlx.Connect("postgres", "user=postgres password=postgres dbname=url_shorter_db sslmode=disable")
	connStr := "host=localhost port=5431 user=postgres password=postgres dbname=url_shorter_db sslmode=disable"
	conn, err := sqlx.Open("postgres", connStr)

	if err != nil {
		// log.Fatalf("Connection to postgres db failed")
		fmt.Println("Db error: ", err)
	}

	return &PostgresDB{conn: conn}, nil
}

type TestSchema struct {
	Id int `db:"id"`
}

func (db *PostgresDB) CreateTables(ctx context.Context) error {
	// fmt.Printf("Tables was created")
	// query := `CREATE TABLE Test(
	// 	id INTEGER
	// );`
	// dbContext, cancel := context.WithTimeout(ctx, 5*time.Second)
	// defer cancel()

	// if cancel == nil {
	// 	log.Fatalf("Db context was not created")
	// }

	// defer cancel()
	// db.conn.QueryRow(dbContext, query)

	// _, err := db.conn.ExecContext(dbContext, query)

	// if err != nil {
	// 	log.Println(err)
	// }

	// _, err := db.conn.Query("INSERT INTO Test VALUES(15)")

	// if err != nil {
	// 	log.Println(err)
	// }

	// rows, err := db.conn.Query("SELECT * FROM Test;")
	// if err != nil {
	// 	log.Println(err)
	// }

	var t TestSchema
	err := db.conn.Get(&t, "SELECT * FROM TEST LIMIT 1")

	if err != nil {
		log.Fatalf("Select query to db was not successs\n %s", err)
	}

	// for rows.Next() {
	// 	err := sqlx.StructScan(rows, &t)

	// 	if err != nil {
	// 		log.Fatalf("Select query to db was not successs\n %s", err)
	// 	}
	// }

	fmt.Println(t)

	return nil
}
