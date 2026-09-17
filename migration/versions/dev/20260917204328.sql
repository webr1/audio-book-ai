-- Create "books" table
CREATE TABLE "public"."books" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "filename" character varying(255) NOT NULL,
  "status" character varying(32) NOT NULL DEFAULT 'pending',
  "word_count" bigint NOT NULL DEFAULT 0,
  "char_count" bigint NOT NULL DEFAULT 0,
  "chunks_total" bigint NOT NULL DEFAULT 0,
  "chunks_done" bigint NOT NULL DEFAULT 0,
  "error_msg" text NOT NULL DEFAULT '',
  "output_path" character varying(512) NOT NULL DEFAULT '',
  PRIMARY KEY ("id")
);
-- Create index "idx_books_deleted_at" to table: "books"
CREATE INDEX "idx_books_deleted_at" ON "public"."books" ("deleted_at");
