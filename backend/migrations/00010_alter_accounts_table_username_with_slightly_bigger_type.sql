-- +goose Up
ALTER TABLE accounts
    ALTER COLUMN username TYPE varchar(30);

-- +goose Down
ALTER TABLE accounts
    ALTER COLUMN username TYPE varchar(8);
