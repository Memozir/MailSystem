-- +goose Up
-- +goose StatementBegin
ALTER TABLE "user" DROP COLUMN phone;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "user" ADD COLUMN phone varchar(12) unique;
-- +goose StatementEnd
