ALTER TABLE users
    ADD COLUMN IF NOT EXISTS is_admin BOOLEAN DEFAULT FALSE;

INSERT INTO users (phone, is_admin)
    VALUES ('+989029266635', TRUE)
        ON CONFLICT (phone) DO UPDATE
            SET is_admin = EXCLUDED.is_admin;
