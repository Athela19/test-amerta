CREATE TABLE IF NOT EXISTS users (
    id            BIGSERIAL PRIMARY KEY,
    username      TEXT NOT NULL UNIQUE CHECK (username = lower(username)),
    password_hash TEXT NOT NULL,
    role          TEXT NOT NULL DEFAULT 'user' CHECK (role IN ('admin', 'user'))
);

CREATE TABLE IF NOT EXISTS refresh_tokens (
    token_hash TEXT PRIMARY KEY,
    user_id    BIGINT NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    expires_at TIMESTAMPTZ NOT NULL,
    revoked_at TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS refresh_tokens_user_idx ON refresh_tokens (user_id);

INSERT INTO users (username, password_hash, role)
VALUES
    ('admin', '$2a$12$5chNCxFg0zvnneYqbXd8w.9hRhT48ekpq0WFHIXvQOaBOuamq/3c6', 'admin'),
    ('user', '$2a$12$zrJxSpywBOzUX.y3sqUntOIXS4jv9kSrs65VQq0SWS6zs9qvdOXHG', 'user')
ON CONFLICT (username) DO NOTHING;
