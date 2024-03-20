CREATE TABLE "products"
(
    "id"             bigserial PRIMARY KEY,
    "name"           varchar,
    "price"          int,
    "stock_quantity" int
);

CREATE TABLE "shelves"
(
    "id"   bigserial PRIMARY KEY,
    "name" varchar
);

CREATE TABLE "products_shelves"
(
    "id"         bigserial PRIMARY KEY,
    "product_id" bigint NOT NULL,
    "shelves_id" bigint NOT NULL,
    "is_primary" bool
);

CREATE TABLE "orders"
(
    "id"         bigserial PRIMARY KEY,
    "order_date" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "order_items"
(
    "id"         bigserial PRIMARY KEY,
    "order_id"   bigint NOT NULL,
    "product_id" bigint NOT NULL,
    "quantity"   int
);

ALTER TABLE "products_shelves"
    ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "products_shelves"
    ADD FOREIGN KEY ("shelves_id") REFERENCES "shelves" ("id");

ALTER TABLE "order_items"
    ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");

ALTER TABLE "order_items"
    ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");