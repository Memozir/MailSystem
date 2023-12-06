-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "role"(
    code integer not null unique primary key,
    "name" varchar(50) not null unique
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "role" IF EXISTS;
-- +goose StatementEnd
