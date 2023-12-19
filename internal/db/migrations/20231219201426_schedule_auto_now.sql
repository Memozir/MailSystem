-- +goose Up
-- +goose StatementBegin
ALTER TABLE delivery_schedule ALTER COLUMN export_date SET default CURRENT_DATE;
ALTER TABLE delivery_schedule ALTER COLUMN import_date SET default CURRENT_DATE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE delivery_schedule ALTER COLUMN export_date DROP default;
ALTER TABLE delivery_schedule ALTER COLUMN import_date DROP default;
-- +goose StatementEnd
