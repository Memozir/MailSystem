-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS employee(
    id bigserial primary key,
    "user" bigint unique,
    "role" bigint,
    foreign key ("user") references "user" ON DELETE CASCADE,
    foreign key ("role") references "role"
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS employee;
-- +goose StatementEnd