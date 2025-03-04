-- +goose Up
-- +goose StatementBegin
CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    category_id INT NOT NULL,
    available BOOLEAN DEFAULT TRUE,
    FOREIGN KEY (category_id) REFERENCES category (id) ON DELETE CASCADE
);
CREATE INDEX idx_books_category_id ON books(category_id);
CREATE INDEX idx_books_title ON books(title);
CREATE INDEX idx_books_author ON books(author);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS books;
-- +goose StatementEnd
