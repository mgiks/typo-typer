-- +goose Up
ALTER TABLE accounts
    ALTER COLUMN username TYPE varchar(8);

ALTER TABLE accounts
    ADD CONSTRAINT not_empty_username_constraint CHECK (length(trim(username)) > 0);

-- +goose Down
ALTER TABLE accounts
    ALTER COLUMN username TYPE varchar(256);

ALTER TABLE accounts
    DROP CONSTRAINT not_empty_username_constraint;
