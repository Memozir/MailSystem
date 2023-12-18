-- +goose Up
-- +goose StatementBegin
ALTER TABLE package ALTER COLUMN "status" SET default 0;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE package ALTER COLUMN "status" DROP default;
-- +goose StatementEnd
