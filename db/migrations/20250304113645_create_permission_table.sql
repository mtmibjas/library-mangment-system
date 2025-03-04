-- +goose Up
-- +goose StatementBegin
CREATE TABLE permissions (
    id SERIAL PRIMARY KEY,
    action VARCHAR(255) NOT NULL UNIQUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS permissions;
-- +goose StatementEnd
