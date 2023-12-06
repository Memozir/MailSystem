-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS delivery_schedule(
    id bigserial primary key,
    courier bigint unique, 
    export_date date,
    import_date date,
    foreign key (courier) references employee
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS delivery_schedule;
-- +goose StatementEnd
