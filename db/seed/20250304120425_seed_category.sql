-- +goose Up
-- +goose StatementBegin
INSERT INTO category (name) VALUES 
    ('Children'),
    ('Teen'),
    ('Adult')
ON CONFLICT (name) DO NOTHING;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM category WHERE name IN ('Children', 'Teen', 'Adult');
-- +goose StatementEnd
