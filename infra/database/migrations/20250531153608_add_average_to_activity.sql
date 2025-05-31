-- +goose Up
-- +goose StatementBegin
ALTER TABLE "activity" ADD COLUMN "average" FLOAT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "activity" DROP COLUMN IF EXISTS "average";
-- +goose StatementEnd
