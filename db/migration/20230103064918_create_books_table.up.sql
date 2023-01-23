-- +migrate Up
CREATE TABLE IF NOT EXISTS "books" (
  "id" BIGINT PRIMARY KEY,
  "title" TEXT NOT NULL,
  "author" TEXT NOT NULL DEFAULT '',
  "description" TEXT NOT NULL DEFAULT '',
  "published_date" DATE,
  "created_at" TIMESTAMP NOT NULL DEFAULT 'now()',
  "updated_at" TIMESTAMP NOT NULL DEFAULT 'now()',
  "deleted_at" TIMESTAMP
);
