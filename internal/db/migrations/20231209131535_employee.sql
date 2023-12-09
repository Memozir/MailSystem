-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS employee(
    id bigserial primary key,
    "user" bigint unique,
    "role" bigint unique,
    foreign key ("user") references "user",
    foreign key ("role") references "role",
    unique("user", "role")
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS employee;
-- +goose StatementEnd