-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS 'address'(
    id bigserial primary key,
    'name' varchar(100) unique not null,
    department references department(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS 'address';
-- +goose StatementEnd
