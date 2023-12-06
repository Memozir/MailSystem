-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS client(
    id bigserial primary key,
    "user" references "user"(id) not null,
    "address" references "address"(id),
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS client;
-- +goose StatementEnd
