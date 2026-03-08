-- +goose Up
CREATE TABLE IF NOT EXISTS refresh_token (
    id bigserial PRIMARY KEY,
    token_hash text NOT NULL,
    salt text NOT NULL,
    account_id uuid REFERENCES account (id),
    created_at timestamp NOT NULL DEFAULT now(),
    expires_at timestamp NOT NULL,
    revoked boolean NOT NULL DEFAULT FALSE
);

-- +goose Down
DROP TABLE IF EXISTS refresh_token;
