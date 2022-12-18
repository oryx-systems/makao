BEGIN;

CREATE TABLE "makao_user" (
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

CREATE TABLE "makao_contact" (
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

CREATE TABLE "makao_user_pin" (
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

CREATE TABLE "makao_user_otp" (
  "id" uuid PRIMARY KEY,
  "created_at" timestamp,
  "created_by" uuid,
  "updated_at" timestamp,
  "updated_by" uuid,
  "is_valid" boolean,
  "valid_until" timestamp NOT NULL,
  "phone_number" varchar(20) NOT NULL,
  "otp" varchar(10) NOT NULL,
  "flavour" varchar(10) NOT NULL,
  "medium" varchar(10),
  "user_id" uuid NOT NULL
);

CREATE TABLE "makao_residence" (
  "id" uuid PRIMARY KEY,
  "created_at" timestamp,
  "created_by" uuid,
  "updated_at" timestamp,
  "updated_by" uuid,
  "active" boolean,
  "name" varchar(100) UNIQUE NOT NULL,
  "registration_number" varchar(50) NOT NULL,
  "location" varchar(100),
  "living_rooms_count" int,
  "owner" uuid NOT NULL
);

CREATE TABLE "makao_identifier" (
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

CREATE TABLE "makao_house" (
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

CREATE TABLE "makao_house_client" (
  "id" uuid PRIMARY KEY,
  "house_id" uuid NOT NULL,
  "tenant_id" uuid NOT NULL
);

CREATE TABLE "makao_bill" (
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

CREATE TABLE "makao_user_residence" (
  "id" uuid PRIMARY KEY,
  "user_id" uuid NOT NULL,
  "residence_id" uuid NOT NULL
);

CREATE UNIQUE INDEX ON "makao_contact" ("flavour", "contact_value");

CREATE UNIQUE INDEX ON "makao_identifier" ("identifier_type", "identifier_value");

CREATE UNIQUE INDEX ON "makao_bill" ("type", "user_id");

CREATE UNIQUE INDEX ON "makao_user_residence" ("user_id", "residence_id");

ALTER TABLE "makao_contact" ADD FOREIGN KEY ("user_id") REFERENCES "makao_user" ("id");

ALTER TABLE "makao_user_pin" ADD FOREIGN KEY ("user_id") REFERENCES "makao_user" ("id");

ALTER TABLE "makao_user_otp" ADD FOREIGN KEY ("user_id") REFERENCES "makao_user" ("id");

ALTER TABLE "makao_residence" ADD FOREIGN KEY ("owner") REFERENCES "makao_user" ("id");

ALTER TABLE "makao_identifier" ADD FOREIGN KEY ("user_id") REFERENCES "makao_user" ("id");

ALTER TABLE "makao_house_client" ADD FOREIGN KEY ("house_id") REFERENCES "makao_house" ("id");

ALTER TABLE "makao_house_client" ADD FOREIGN KEY ("tenant_id") REFERENCES "makao_user" ("id");

ALTER TABLE "makao_bill" ADD FOREIGN KEY ("user_id") REFERENCES "makao_user" ("id");

ALTER TABLE "makao_user_residence" ADD FOREIGN KEY ("user_id") REFERENCES "makao_user" ("id");

ALTER TABLE "makao_user_residence" ADD FOREIGN KEY ("residence_id") REFERENCES "makao_residence" ("id");

COMMIT;