-- +goose Up
CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS account (
    id uuid DEFAULT uuidv7 () PRIMARY KEY,
    username varchar(256) UNIQUE NOT NULL,
    email citext UNIQUE,
    passhash text NOT NULL,
    salt text NOT NULL,
    wpm double precision
);

-- +goose Down
DROP TABLE IF EXISTS account;
