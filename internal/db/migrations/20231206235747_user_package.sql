-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_package(
    "user" references "user"(id),
    package references package(id),
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_package;
-- +goose StatementEnd
