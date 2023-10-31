migrate:
	goose postgres "user=postgres password=postgres host=localhost port=5431 dbname=mail_system_db sslmode=disable" up