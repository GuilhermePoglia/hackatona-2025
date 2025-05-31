-- +goose Up
-- +goose StatementBegin
ALTER TABLE "activity" 
ADD COLUMN "image" VARCHAR(500);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "activity" 
DROP COLUMN IF EXISTS "image";
-- +goose StatementEnd
