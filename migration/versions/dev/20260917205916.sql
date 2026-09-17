-- Modify "books" table
ALTER TABLE "public"."books" ADD COLUMN "user_id" bigint NOT NULL;
-- Create index "idx_books_user_id" to table: "books"
CREATE INDEX "idx_books_user_id" ON "public"."books" ("user_id");
-- Create "users" table
CREATE TABLE "public"."users" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "google_sub" character varying(255) NOT NULL,
  "email" character varying(255) NOT NULL,
  "name" character varying(255) NOT NULL DEFAULT '',
  "picture" character varying(512) NOT NULL DEFAULT '',
  PRIMARY KEY ("id")
);
-- Create index "idx_users_deleted_at" to table: "users"
CREATE INDEX "idx_users_deleted_at" ON "public"."users" ("deleted_at");
-- Create index "idx_users_google_sub" to table: "users"
CREATE UNIQUE INDEX "idx_users_google_sub" ON "public"."users" ("google_sub");
