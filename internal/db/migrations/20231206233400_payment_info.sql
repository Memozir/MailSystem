-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS payment_info(
    id serial primary key,
    package bigint not null unique,
    tarrif bigint not null,
    foreign key (package) references package,
    foregin key (tarrif) references tarrif
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS payment_info;
-- +goose StatementEnd
