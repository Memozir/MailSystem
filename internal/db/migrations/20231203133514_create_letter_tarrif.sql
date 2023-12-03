-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS letter_tariff(
    price bigint unique,
    'weight' bigint unique
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
