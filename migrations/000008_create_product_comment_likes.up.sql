CREATE TABLE IF NOT EXISTS product_comment_likes
(
    comment_id INTEGER  NOT NULL REFERENCES product_comments (id) ON DELETE CASCADE,
    user_id    INTEGER  NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    vote       SMALLINT NOT NULL CHECK (vote IN (-1, 1)),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (comment_id, user_id)
);

CREATE INDEX IF NOT EXISTS idx_product_comment_likes_comment_id ON product_comment_likes (comment_id);
CREATE INDEX IF NOT EXISTS idx_product_comment_likes_user_id ON product_comment_likes (user_id);
