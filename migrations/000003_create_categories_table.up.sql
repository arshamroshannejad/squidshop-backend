CREATE TABLE IF NOT EXISTS categories
(
    id        SERIAL PRIMARY KEY,
    name      VARCHAR(100) NOT NULL UNIQUE,
    slug      VARCHAR(100) NOT NULL UNIQUE,
    parent_id INT DEFAULT NULL,
    FOREIGN KEY (parent_id) REFERENCES categories (id) ON DELETE CASCADE
);

CREATE INDEX idx_categories_parent_id ON categories (parent_id);
ALTER TABLE categories
    ADD CONSTRAINT check_parent_not_self CHECK (id != parent_id);
