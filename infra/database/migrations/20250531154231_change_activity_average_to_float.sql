-- +goose Up
-- +goose StatementBegin
ALTER TABLE activity ALTER COLUMN average TYPE REAL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE activity ALTER COLUMN average TYPE DECIMAL(3,2);
-- +goose StatementEnd
