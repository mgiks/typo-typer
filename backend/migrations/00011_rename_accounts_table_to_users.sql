-- +goose Up
ALTER TABLE accounts RENAME TO users;

-- +goose Down
ALTER TABLE users RENAME TO accounts;
