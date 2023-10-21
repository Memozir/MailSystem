-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Packages(
    id SERIAL PRIMARY KEY,
    [status] INTEGER,
    [weight] INTEGER NOT NULL,
    price INTEGER NOT NULL,
    sender INTEGER,
    consumer INTEGER,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    delivering_at TIMESTAMP NOT NULL,
    FOREIGN KEY ([status]) REFERENCES Statuses(id)
    FOREIGN KEY (sender) REFERENCES User(id),
    FOREIGN KEY (consumer) REFERENCES Users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Packages;
-- +goose StatementEnd
