-- +goose Up
-- +goose StatementBegin
ALTER TABLE package DROP COLUMN courier;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE package ADD COLUMN courier bigint;
ALTER TABLE package ADD FOREIGN KEY (courier) references employee;
-- +goose StatementEnd
