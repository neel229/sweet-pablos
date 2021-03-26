CREATE TABLE "uaccount" (
  "id" bigserial PRIMARY KEY,
  "first_name" varchar(32) NOT NULL,
  "last_name" varchar(32) NOT NULL,
  "email" varchar(32) NOT NULL,
  "password" varchar(128) NOT NULL,
  "ts" timestamptz NOT NULL DEFAULT (now())
);
CREATE INDEX ON "uaccount" ("email");