-- +goose Up
-- +goose StatementBegin
ALTER TABLE address DROP COLUMN IF EXISTS apartment;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE address ADD COLUMN IF NOT EXISTS apartment varchar(30) not null default '';
-- +goose StatementEnd
