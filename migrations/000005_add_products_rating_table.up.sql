CREATE TABLE IF NOT EXISTS product_ratings
(
    product_id INT NOT NULL REFERENCES products (id) ON DELETE CASCADE,
    user_id    INT NOT NULL REFERENCES users (id) ON DELETE SET NULL,
    rating     INT NOT NULL CHECK (rating >= 1 AND rating <= 5),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (product_id, user_id)
);

CREATE INDEX IF NOT EXISTS idx_product_ratings_product_id ON product_ratings (product_id);
CREATE INDEX IF NOT EXISTS idx_product_ratings_user_id ON product_ratings (user_id);
CREATE INDEX IF NOT EXISTS idx_product_ratings_rating ON product_ratings (rating);
