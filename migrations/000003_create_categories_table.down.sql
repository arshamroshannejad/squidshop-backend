ALTER TABLE IF EXISTS categories DROP CONSTRAINT IF EXISTS check_parent_not_self;
DROP INDEX IF EXISTS idx_categories_parent_id;
DROP TABLE IF EXISTS categories CASCADE;
