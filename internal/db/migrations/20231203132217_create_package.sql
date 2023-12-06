-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS package(
    id bigserial primary key,
    "status" int,
    "weight" bigint,
    sender references client(id) not null,
    receiver references client(id) not null,
    create_date date,
    deliver_date date,
    curier references employee(id),
    "type" smallint not null,
    department_receiver references department(id),
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS package;
-- +goose StatementEnd
