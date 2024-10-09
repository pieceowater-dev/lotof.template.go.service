-- Create "items" table
CREATE TABLE "public"."items" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "name" text NULL,
  "comment" text NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_items_deleted_at" to table: "items"
CREATE INDEX "idx_items_deleted_at" ON "public"."items" ("deleted_at");
