-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS storehouse(
    id serial primary key,
    department references department(id),
    package references package(id),
    is_import boolean,
    is_export boolean
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS storehouse;
-- +goose StatementEnd
