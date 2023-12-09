-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS package(
    id bigserial primary key,
    "status" int,
    "weight" bigint,
    sender bigint,
    receiver bigint,
    courier bigint,
    department_receiver bigint,
    "type" smallint not null,
    create_date date,
    deliver_date date,
    foreign key (sender) references client,
    foreign key (receiver) references client,
    foreign key (courier) references employee,
    foreign key (department_receiver) references department
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS package;
-- +goose StatementEnd