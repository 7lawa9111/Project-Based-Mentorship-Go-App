CREATE TABLE "authors" (
                           "id" UUID PRIMARY KEY,
                           "first_name" VARCHAR,
                           "last_name" VARCHAR,
                           "created_at" timestamp,
                           "updated_at" timestamp,
                           "created_by" VARCHAR,
                           "updated_by" VARCHAR
);

CREATE TABLE "documents" (
                             "id" UUID PRIMARY KEY,
                             "author_id" UUID,
                             "title" VARCHAR,
                             "content" BYTEA,
                             "created_at" timestamp,
                             "updated_at" timestamp,
                             "created_by" VARCHAR,
                             "updated_by" VARCHAR
);

ALTER TABLE "documents" ADD CONSTRAINT "documents" FOREIGN KEY ("author_id") REFERENCES "authors" ("id");