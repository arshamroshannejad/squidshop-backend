CREATE TABLE IF NOT EXISTS users
(
    id         SERIAL PRIMARY KEY,
    phone      VARCHAR(20) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_users_phone ON users (phone);
CREATE INDEX IF NOT EXISTS idx_users_created_at ON users (created_at);
