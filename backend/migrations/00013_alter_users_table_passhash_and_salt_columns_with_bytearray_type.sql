-- +goose Up
ALTER TABLE users
    ALTER COLUMN passhash TYPE bytea
    USING convert_to(passhash, 'UTF8');

ALTER TABLE users
    ALTER COLUMN salt TYPE bytea
    USING convert_to(salt, 'UTF8');

-- +goose Down
ALTER TABLE users
    ALTER COLUMN passhash TYPE text
    USING convert_from(passhash, 'UTF8');

ALTER TABLE users
    ALTER COLUMN salt TYPE text
    USING convert_from(salt, 'UTF8');
