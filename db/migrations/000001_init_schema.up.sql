
CREATE TABLE "User" (
  "id" uuid PRIMARY KEY,
  "name" varchar NOT NULL,
  "profile_pic" varchar
);

CREATE TABLE "UserInterests" (
  "user_id" uuid,
  "interest_id" uuid,
  PRIMARY KEY ("user_id", "interest_id")
);

-- I had to add a table for Categories because UserInterests references it
CREATE TABLE "Category" (
  "id" uuid PRIMARY KEY,
  "category" varchar,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  "deleted_at" timestamp DEFAULT null
);

ALTER TABLE "UserInterests" ADD FOREIGN KEY ("user_id") REFERENCES "User" ("id");

ALTER TABLE "UserInterests" ADD FOREIGN KEY ("interest_id") REFERENCES "Category" ("id");