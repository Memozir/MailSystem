-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS client(
    id bigserial primary key,
    "user" bigint not null,
    "address" bigint,
    foreign key ("user") references "user",
    foreign key ("address") references "address"
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS client;
-- +goose StatementEnd
