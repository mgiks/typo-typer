-- +goose Up
ALTER TABLE refresh_tokens
    ALTER COLUMN id DROP DEFAULT;

ALTER SEQUENCE refresh_tokens_id_seq
    RENAME TO refresh_tokens_id_seq_bck;

ALTER TABLE refresh_tokens
    ALTER COLUMN id
    ADD GENERATED ALWAYS AS IDENTITY;

SELECT
    setval('public.refresh_tokens_id_seq', GREATEST (COALESCE((
                SELECT
                    MAX(id)
                FROM refresh_tokens), 0), 1), TRUE);

DROP SEQUENCE refresh_tokens_id_seq_bck;

-- +goose Down
CREATE SEQUENCE refresh_tokens_id_seq_tmp OWNED BY refresh_tokens.id;

SELECT
    setval('public.refresh_tokens_id_seq_tmp', GREATEST (COALESCE((
                SELECT
                    MAX(id)
                FROM refresh_tokens), 0), 1), TRUE);

ALTER TABLE refresh_tokens
    ALTER COLUMN id
    DROP IDENTITY;

ALTER SEQUENCE refresh_tokens_id_seq_tmp
    RENAME TO refresh_tokens_id_seq;

ALTER TABLE refresh_tokens
    ALTER COLUMN id SET DEFAULT nextval('refresh_tokens_id_seq');
