-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXITS delivery_schedule(
    id bigserial primary key,
    courier references employee(id) unique,
    export_date date,
    import_date date 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS delivery_schedule;
-- +goose StatementEnd
