-- +goose Up
ALTER TABLE texts RENAME COLUMN entry TO content;

-- +goose Down
ALTER TABLE texts RENAME COLUMN content TO entry;
