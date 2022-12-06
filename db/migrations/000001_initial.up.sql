BEGIN;

CREATE TABLE "user" (
  "id" uuid PRIMARY KEY,
  "created_at" timestamp,
  "created_by" uuid,
  "updated_at" timestamp,
  "updated_by" uuid,
  "first_name" varchar(20) NOT NULL,
  "middle_name" varchar(20),
  "last_name" varchar(20) NOT NULL,
  "active" boolean,
  "flavour" varchar(10) NOT NULL,
  "username" varchar(20) UNIQUE NOT NULL,
  "user_type" varchar(20),
  "device_token" text,
  "residence" uuid
);

CREATE TABLE "contact" (
  "id" uuid PRIMARY KEY,
  "created_at" timestamp,
  "created_by" uuid,
  "updated_at" timestamp,
  "updated_by" uuid,
  "active" boolean,
  "contact_type" varchar(10),
  "contact_value" varchar(20),
  "flavour" varchar(10) NOT NULL,
  "user_id" uuid NOT NULL
);

CREATE TABLE "user_pin" (
  "id" uuid PRIMARY KEY,
  "created_at" timestamp,
  "created_by" uuid,
  "updated_at" timestamp,
  "updated_by" uuid,
  "active" boolean,
  "flavour" varchar(10),
  "valid_from" timestamp NOT NULL,
  "valid_to" timestamp NOT NULL,
  "hashed_pin" text NOT NULL,
  "salt" text NOT NULL,
  "user_id" uuid NOT NULL
);

CREATE TABLE "user_otp" (
  "id" uuid PRIMARY KEY,
  "created_at" timestamp,
  "created_by" uuid,
  "updated_at" timestamp,
  "updated_by" uuid,
  "is_valid" boolean,
  "valid_until" timestamp NOT NULL,
  "phonenumber" varchar(20) NOT NULL,
  "otp" varchar(10) NOT NULL,
  "flavour" varchar(10) NOT NULL,
  "medium" varchar(10),
  "user_id" uuid NOT NULL
);

CREATE TABLE "residence" (
  "id" uuid PRIMARY KEY,
  "created_at" timestamp,
  "created_by" uuid,
  "updated_at" timestamp,
  "updated_by" uuid,
  "active" boolean,
  "name" varchar(50) NOT NULL,
  "registration_number" varchar(50) NOT NULL,
  "location" varchar(100),
  "living_rooms_count" int,
  "owner" uuid NOT NULL
);

CREATE TABLE "identifier" (
  "id" uuid PRIMARY KEY,
  "created_at" timestamp,
  "created_by" uuid,
  "updated_at" timestamp,
  "updated_by" uuid,
  "active" boolean,
  "identifier_type" varchar(20),
  "identifier_value" varchar(20),
  "user_id" uuid NOT NULL
);

CREATE TABLE "house" (
  "id" uuid PRIMARY KEY,
  "created_at" timestamp,
  "created_by" uuid,
  "updated_at" timestamp,
  "updated_by" uuid,
  "active" boolean,
  "house_number" varchar(20) UNIQUE NOT NULL,
  "house_category" varchar(20) NOT NULL,
  "house_class" varchar(20),
  "rent_value" float NOT NULL
);

CREATE TABLE "house_client" (
  "id" uuid PRIMARY KEY,
  "house_id" uuid NOT NULL,
  "tenant_id" uuid NOT NULL
);

CREATE TABLE "bill" (
  "id" uuid PRIMARY KEY,
  "created_at" timestamp,
  "created_by" uuid,
  "updated_at" timestamp,
  "updated_by" uuid,
  "active" boolean,
  "type" varchar(20),
  "amount" float DEFAULT 0,
  "penalty" float DEFAULT 0,
  "user_id" uuid NOT NULL
);

CREATE UNIQUE INDEX ON "identifier" ("identifier_type", "identifier_value");

CREATE UNIQUE INDEX ON "bill" ("type", "user_id");

ALTER TABLE "contact" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "user_pin" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "user_otp" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "residence" ADD FOREIGN KEY ("owner") REFERENCES "user" ("id");

ALTER TABLE "identifier" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "house_client" ADD FOREIGN KEY ("house_id") REFERENCES "house" ("id");

ALTER TABLE "house_client" ADD FOREIGN KEY ("tenant_id") REFERENCES "user" ("id");

ALTER TABLE "bill" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

COMMIT;