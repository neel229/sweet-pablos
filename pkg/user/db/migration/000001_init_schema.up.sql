CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "first_name" varchar(32) NOT NULL,
  "last_name" varchar(32),
  "email" varchar(32) NOT NULL,
  "password" varchar(32) NOT NULL,
  "ts" timestamptz NOT NULL DEFAULT (now())
);
CREATE INDEX ON "users" ("email");