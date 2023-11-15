

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

    
-- -- Employee
-- -- Category
-- -- Product

-- -- Remaining
--     Category
--     Product
--     Count
--     Price
--     Total Price

-- -- Sales
--     sale_id -> S-0000001
--     address
--     datetime
--     Total Sum
--     Total Count
--     -- Sales_Product
--         Category
--         Product
--         Count
--         Price
--         Total Price



