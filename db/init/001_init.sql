CREATE TABLE IF NOT EXISTS users (
    id            BIGSERIAL PRIMARY KEY,
    email         TEXT NOT NULL UNIQUE CHECK (email = lower(email)),
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

INSERT INTO users (email, password_hash, role)
VALUES
    ('admin@amerta.local', '$2a$12$LwrMfRUySqihL.XHT663IugmqeCCbzhm.pOBaaXMPaj8us/LixRJe', 'admin'),
    ('user@amerta.local', '$2a$12$3.ClLVysPTx8k4tT/qFYrecfrkkkFA7/WgIl9p3DemZUyfwTMd6Ya', 'user')
ON CONFLICT (email) DO NOTHING;
