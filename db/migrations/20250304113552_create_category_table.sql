-- +goose Up
-- +goose StatementBegin
CREATE TABLE category (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE
);
CREATE INDEX idx_category_name ON category(name);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS category;
-- +goose StatementEnd
