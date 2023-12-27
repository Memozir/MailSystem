-- +goose Up
-- +goose StatementBegin
ALTER TABLE "client" ADD COLUMN apartment varchar(30) not null default '';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "client" DROP COLUMN apartment;
-- +goose StatementEnd
