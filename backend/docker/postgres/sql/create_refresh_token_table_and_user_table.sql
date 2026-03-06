CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS account (
    id uuid DEFAULT uuidv7 () PRIMARY KEY,
    username varchar(256) UNIQUE NOT NULL,
    email citext UNIQUE,
    passhash text NOT NULL,
    salt text NOT NULL,
    wpm double precision
);

CREATE TABLE IF NOT EXISTS refresh_token (
    id bigserial PRIMARY KEY,
    token_hash text NOT NULL,
    salt text NOT NULL,
    account_id uuid REFERENCES account (id),
    created_at timestamp NOT NULL DEFAULT now(),
    expires_at timestamp NOT NULL,
    revoked boolean NOT NULL DEFAULT FALSE
);
