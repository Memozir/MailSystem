-- +goose Up
-- +goose StatementBegin
ALTER TABLE "address" ADD COLUMN apartment varchar(30) not null default '';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "address" DROP COLUMN apartment;
-- +goose StatementEnd
