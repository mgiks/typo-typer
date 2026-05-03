-- +goose Up
ALTER TABLE typing_text RENAME TO texts;

ALTER TABLE account RENAME TO accounts;

ALTER TABLE refresh_token RENAME TO refresh_tokens;

-- +goose Down
ALTER TABLE texts RENAME TO typing_text;

ALTER TABLE accounts RENAME TO account;

ALTER TABLE refresh_tokens RENAME TO refresh_token;
