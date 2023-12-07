-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_package(
    "user" bigint not null,
    package bigint not null,
    foreign key ("user") references "user",
    foreign key (package) references package
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_package;
-- +goose StatementEnd
