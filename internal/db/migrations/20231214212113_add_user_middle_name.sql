-- +goose Up
-- +goose StatementBegin
ALTER TABLE "user" ADD COLUMN middle_name varchar(50) unique;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "user" DROP COLUMN middle_name;
-- +goose StatementEnd
