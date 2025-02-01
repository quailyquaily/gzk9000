-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TABLE IF NOT EXISTS "properties" (
  "key" character varying(64) PRIMARY KEY,
  "value" character varying(1024),
  "updated_at" timestamp with time zone
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS "properties";
-- +goose StatementEnd
