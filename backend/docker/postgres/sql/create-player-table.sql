CREATE TABLE IF NOT EXISTS "player" (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name varchar(128) NOT NULL,
    email varchar(320) NOT NULL,
    hashed_password varchar(256) NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);
