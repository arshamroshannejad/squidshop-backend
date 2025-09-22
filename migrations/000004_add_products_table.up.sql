CREATE TABLE IF NOT EXISTS products
(
    id                SERIAL PRIMARY KEY,
    name              VARCHAR(255)   NOT NULL,
    slug              VARCHAR(255)   NOT NULL UNIQUE,
    description       TEXT,
    short_description VARCHAR(255),
    price             DECIMAL(10, 2) NOT NULL,
    quantity          INTEGER        NOT NULL DEFAULT 0,
    created_at        TIMESTAMP               DEFAULT CURRENT_TIMESTAMP,
    updated_at        TIMESTAMP               DEFAULT CURRENT_TIMESTAMP,
    category_id       INTEGER        REFERENCES categories (id) ON DELETE SET NULL,
    CHECK (price >= 0),
    CHECK (quantity >= 0)
);

CREATE INDEX idx_products_name ON products (name);
CREATE INDEX idx_products_price ON products (price);
CREATE INDEX idx_products_quantity ON products (quantity);
CREATE INDEX idx_products_category ON products (category_id);
CREATE INDEX idx_products_created_at ON products (created_at);
CREATE INDEX idx_products_search ON products USING GIN (to_tsvector('english', name || ' ' || COALESCE(description, '')));
