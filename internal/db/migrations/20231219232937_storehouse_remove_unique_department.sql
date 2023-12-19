-- +goose Up
-- +goose StatementBegin
ALTER TABLE storehouse DROP COLUMN department;
ALTER TABLE storehouse ADD COLUMN department bigint;
ALTER TABLE storehouse ADD FOREIGN KEY (department) references department;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
