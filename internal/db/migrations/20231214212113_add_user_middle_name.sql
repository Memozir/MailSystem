-- +goose Up
-- +goose StatementBegin
ALTER TABLE "user" ADD COLUMN IF NOT EXISTS middle_name varchar(50);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "user" DROP COLUMN IF EXISTS middle_name;
-- +goose StatementEnd
