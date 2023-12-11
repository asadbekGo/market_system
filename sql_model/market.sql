

CREATE TABLE "category" (
    "id" UUID NOT NULL PRIMARY KEY,
    "title" VARCHAR(46) NOT NULL,
    "parent_id" UUID REFERENCES "category" ("id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE "product" (
    "id" UUID NOT NULL PRIMARY KEY,
    "name" VARCHAR(46) NOT NULL,
    "barcode" VARCHAR NOT NULL,
    "price" NUMERIC NOT NULL,
    "image_url" VARCHAR(255) NOT NULL,
    "category_id" UUID NOT NULL REFERENCES "category"("id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE "user" (
    "id" UUID NOT NULL PRIMARY KEY,
    "first_name" VARCHAR(46) NOT NULL,
    "last_name" VARCHAR(46) NOT NULL,
    "login" VARCHAR(46) NOT NULL,
    "password" VARCHAR NOT NULL,
    "active" BOOLEAN NOT NULL DEFAULT true,
    "client_type" VARCHAR(46) NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

