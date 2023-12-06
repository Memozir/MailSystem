-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS payment_info(
    id serial primary key,
    package references package(id),
    tarrif references tarrif(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS payment_info;
-- +goose StatementEnd
