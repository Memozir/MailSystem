-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tarrif(
    id primary key,
    price bigint,
    "weight" bigint,
    "type" smallint,
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tarrif IF EXISTS;
-- +goose StatementEnd
