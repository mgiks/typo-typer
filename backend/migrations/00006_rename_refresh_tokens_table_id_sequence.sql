-- +goose Up
ALTER SEQUENCE refresh_token_id_seq
    RENAME TO refresh_tokens_id_seq;

-- +goose Down
ALTER SEQUENCE refresh_tokens_id_seq
    RENAME TO refresh_token_id_seq;
