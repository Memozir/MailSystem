-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "user"(
    id bigserial primary key,
    pass varchar(20) not null,
    "login" varchar(30) not null UNIQUE,
    birth_date date,
    first_name varchar(30),
    last_name varchar(30)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "user";
-- +goose StatementEnd