-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tarrif (
    id bigserial primary key,
    price bigint,
    "weight" bigint,
    "type" smallint
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tarrif;
-- +goose StatementEnd
