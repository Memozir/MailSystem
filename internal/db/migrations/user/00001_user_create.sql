-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Users(
    id SERIAL PRIMARY KEY,
    phone VARCHAR(11) NOT NULL,
    pass VARCHAR(20) NOT NULL,
    first_name VARCHAR(50) NOT NULL,
    second_name VARCHAR(50) NOT NULL,
    birth_date DATE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Users;
-- +goose StatementEnd
