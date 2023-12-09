-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS storehouse(
    id serial primary key,
    department bigint unique,
    package bigint not null,
    is_import boolean,
    is_export boolean,
    foreign key (department) references department,
    foreign key (package) references package 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS storehouse;
-- +goose StatementEnd