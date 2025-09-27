CREATE TABLE IF NOT EXISTS product_comments
(
    id         SERIAL PRIMARY KEY,
    product_id INTEGER NOT NULL REFERENCES products (id) ON DELETE CASCADE,
    user_id    INTEGER NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    parent_id  INTEGER REFERENCES product_comments (id) ON DELETE CASCADE,
    comment    TEXT    NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_product_comments_product_id ON product_comments (product_id);
CREATE INDEX IF NOT EXISTS idx_product_comments_user_id ON product_comments (user_id);
CREATE INDEX IF NOT EXISTS idx_product_comments_parent_id ON product_comments (parent_id);
CREATE INDEX IF NOT EXISTS idx_product_comments_created_at ON product_comments (created_at);
