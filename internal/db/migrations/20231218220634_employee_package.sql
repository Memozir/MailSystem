-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS employee_package(
     employee bigint not null,
     package bigint not null,
     foreign key (employee) references employee,
     foreign key (package) references package
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS employee_package;
-- +goose StatementEnd
