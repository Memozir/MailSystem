-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS department(
    id serial primary key,
    is_central boolean
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS department;
-- +goose StatementEnd
