-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TABLE IF NOT EXISTS "agents" (
  "id" serial PRIMARY KEY,
  "name" varchar(128) NOT NULL,
  "display_name" varchar(512) NOT NULL,

  "status" INT NOT NULL DEFAULT 0,

  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now()
);


CREATE TABLE IF NOT EXISTS "facts" (
  "id" serial PRIMARY KEY,
  "agent_id" BIGINT NOT NULL,
  "content" text NOT NULL DEFAULT '',
  "sentiment" DECIMAL(5, 3) NOT NULL DEFAULT 0.0,

  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now()
);


CREATE TABLE IF NOT EXISTS "memslices" (
  "id" serial PRIMARY KEY,
  "response_type" INT NOT NULL,

  "agent_id" BIGINT NOT NULL,
  "speaker_id" BIGINT NOT NULL,

  "is_monolog" boolean NOT NULL DEFAULT false,

  "included_fact_ids" BIGINT[] NOT NULL DEFAULT '{}',
  "external_fact_ids" BIGINT[] NOT NULL DEFAULT '{}',

  "related_memslice_ids" BIGINT[] NOT NULL DEFAULT '{}',

  "content" text NOT NULL,

  "status" INT NOT NULL DEFAULT 0,

  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now()
);


CREATE TABLE IF NOT EXISTS "studygoals" (
  "id" serial PRIMARY KEY,
  "agent_id" BIGINT NOT NULL,
  "content" text NOT NULL DEFAULT '',

  "iteration" INT NOT NULL DEFAULT 0,

  "priority_score" DECIMAL(64, 8) NOT NULL DEFAULT 0.0,

  "status" INT NOT NULL DEFAULT 0,

  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS "studygoals";
DROP TABLE IF EXISTS "memslices";
DROP TABLE IF EXISTS "facts";
DROP TABLE IF EXISTS "agents";

-- +goose StatementEnd
