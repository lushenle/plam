CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "role" varchar NOT NULL DEFAULT 'user',
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01',
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT '0001-01-01'
);

CREATE TABLE "project" (
   "id" char DEFAULT uuid_generate_v4() NOT NULL,
   "name" varchar NOT NULL,
   "amount" float4 NOT NULL,
   "description" text NOT NULL,
   "created_at" timestamptz NOT NULL DEFAULT NOW(),
   "updated_at" timestamptz NOT NULL DEFAULT '0001-01-01',
   PRIMARY KEY ("id")
);

CREATE TABLE "income" (
   "id" char DEFAULT uuid_generate_v4() NOT NULL,
   "payee" varchar NOT NULL,
   "amount" float4 NOT NULL,
   "project_id" uuid NOT NULL,
   "created_at" timestamptz NOT NULL DEFAULT NOW(),
   "updated_at" timestamptz NOT NULL DEFAULT '0001-01-01',
   PRIMARY KEY ("id"),
   FOREIGN KEY ("project_id") REFERENCES "project" ("id")
);

CREATE TABLE "loan" (
   "id" char DEFAULT uuid_generate_v4() NOT NULL,
   "borrower" varchar NOT NULL,
   "amount" float4 NOT NULL,
   "subject" text NOT NULL,
   "created_at" timestamptz NOT NULL DEFAULT NOW(),
   "updated_at" timestamptz NOT NULL DEFAULT '0001-01-01',
   PRIMARY KEY ("id")
);

CREATE TABLE "pay_out" (
   "id" char DEFAULT uuid_generate_v4() NOT NULL,
   "owner" varchar NOT NULL,
   "amount" float4 NOT NULL,
   "subject" text NOT NULL,
   "created_at" timestamptz NOT NULL DEFAULT NOW(),
   "updated_at" timestamptz NOT NULL DEFAULT '0001-01-01',
   PRIMARY KEY ("id")
);

CREATE OR REPLACE FUNCTION update_modified_column ()
   RETURNS TRIGGER
   AS $$
BEGIN
   NEW.updated_at = NOW();
   RETURN NEW;
END;
$$
LANGUAGE 'plpgsql';

CREATE TRIGGER update_income_modtime
   BEFORE UPDATE ON "income"
   FOR EACH ROW
   EXECUTE PROCEDURE update_modified_column ();

CREATE TRIGGER update_project_modtime
   BEFORE UPDATE ON "project"
   FOR EACH ROW
   EXECUTE PROCEDURE update_modified_column ();

CREATE TRIGGER update_loan_modtime
   BEFORE UPDATE ON "loan"
   FOR EACH ROW
   EXECUTE PROCEDURE update_modified_column ();

CREATE TRIGGER update_pay_out_modtime
   BEFORE UPDATE ON "pay_out"
   FOR EACH ROW
   EXECUTE PROCEDURE update_modified_column ();
