-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS user_package;
CREATE TABLE IF NOT EXISTS client_package(
    client bigint not null,
    package bigint not null,
    foreign key (client) references client,
    foreign key (package) references package
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS client_package;
-- +goose StatementEnd
