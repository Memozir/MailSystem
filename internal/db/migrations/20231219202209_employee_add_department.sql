-- +goose Up
-- +goose StatementBegin
ALTER TABLE employee ADD COLUMN department bigint NULL;
ALTER TABLE employee ADD FOREIGN KEY (department) REFERENCES department;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE employee DROP COLUMN department;
-- +goose StatementEnd
