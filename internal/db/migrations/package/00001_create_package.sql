-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Packages(
    id SERIAL PRIMARY KEY,
    "status" INTEGER REFERENCES Statuses(id),
    "weight" INTEGER NOT NULL,
    price INTEGER NOT NULL,
    sender BIGINT REFERENCES Users(id),
    consumer BIGINT REFERENCES Users(id),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    delivering_at TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Packages;
-- +goose StatementEnd
