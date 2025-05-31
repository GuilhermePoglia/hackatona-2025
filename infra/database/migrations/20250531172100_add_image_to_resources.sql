-- +goose Up
-- +goose StatementBegin
ALTER TABLE "resource" 
ADD COLUMN "image" VARCHAR(500);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "resource" 
DROP COLUMN IF EXISTS "image";
-- +goose StatementEnd
