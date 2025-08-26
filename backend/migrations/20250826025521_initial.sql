-- Create "users" table
CREATE TABLE "users" (
  "id" uuid NOT NULL DEFAULT uuidv7(),
  "firebase_uid" character varying(128) NOT NULL,
  "email" character varying(255) NOT NULL,
  "name" character varying(255) NOT NULL,
  "profile_image_url" text NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY ("id"),
  CONSTRAINT "users_firebase_uid_key" UNIQUE ("firebase_uid"),
  CONSTRAINT "users_email_check" CHECK ((email)::text ~ '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$'::text),
  CONSTRAINT "users_name_check" CHECK (length(TRIM(BOTH FROM name)) > 0)
);
-- Create index "idx_users_created_at" to table: "users"
CREATE INDEX "idx_users_created_at" ON "users" ("created_at");
-- Create index "idx_users_email" to table: "users"
CREATE INDEX "idx_users_email" ON "users" ("email");
-- Create index "idx_users_firebase_uid" to table: "users"
CREATE INDEX "idx_users_firebase_uid" ON "users" ("firebase_uid");
