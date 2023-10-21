-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Statuses(
    id SERIAL PRIMARY KEY,
    title VARCHAR(30) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Statuses;
-- +goose StatementEnd
