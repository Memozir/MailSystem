-- +goose Up
-- +goose StatementBegin
ALTER TABLE "user" ADD COLUMN "login" varchar(30) NOT NULL UNIQUE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "user" DROP COLUMN "login";
-- +goose StatementEnd
