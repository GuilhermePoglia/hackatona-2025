-- +goose Up
-- +goose StatementBegin
CREATE TABLE "resource" (
    "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "name" VARCHAR(255),
    "type" VARCHAR(50),
    "midia" TEXT,
    "average" FLOAT,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "resource";
-- +goose StatementEnd
