-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS employee(
    id bigserial primary key,
    user references user(id) unique,
    'role' references 'role'(code) unique
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS employee;
-- +goose StatementEnd
