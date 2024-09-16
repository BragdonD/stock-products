CREATE TABLE "Products" (
    "uuid" uuid NOT NULL,
    "name" varchar(255) NOT NULL,
    "reference" varchar(255) NOT NULL,
    "line" varchar(255) NOT NULL,
    "supplier" varchar(255) NOT NULL,
    "description" text NOT NULL,
    "image-url" varchar(255) NOT NULL,
    "price" decimal(10,2) NOT NULL,
    "created-at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated-at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("uuid")
);

CREATE TABLE "ProductVariant" (
    "uuid" uuid NOT NULL,
    "product-uuid" uuid NOT NULL,
    "name" varchar(255) NOT NULL,
    "reference" varchar(255) NOT NULL,
    "description" text NOT NULL,
    "image-url" varchar(255) NOT NULL,
    "price" decimal(10,2) NOT NULL,
    "created-at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated-at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("uuid"),
    FOREIGN KEY ("product-uuid") REFERENCES "Products"("uuid")
);