-- +goose Up
-- +goose StatementBegin
CREATE TABLE "employee" (
    "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "name" VARCHAR(255),
    "email" VARCHAR(255),
    "position" VARCHAR(255),
    "balance" FLOAT,
    "average" FLOAT,
    "qrcode" VARCHAR(255),
    "password_hash" VARCHAR(255),
    "midia" TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "employee";
-- +goose StatementEnd