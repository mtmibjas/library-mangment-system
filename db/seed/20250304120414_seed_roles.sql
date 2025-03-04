-- +goose Up
-- +goose StatementBegin
INSERT INTO roles (name) VALUES 
    ('admin'),
    ('member'),
    ('librarian')
ON CONFLICT (name) DO NOTHING;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM roles WHERE name IN ('admin', 'member', 'librarian');
-- +goose StatementEnd
