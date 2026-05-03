-- +goose Up
ALTER TABLE texts
    ALTER COLUMN id TYPE bigint;

-- +goose Down
ALTER TABLE texts
    ALTER COLUMN id TYPE int;
