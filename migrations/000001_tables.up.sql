CREATE TABLE products(
    product_id  UUID      PRIMARY KEY       NOT NULL,
    name        VARCHAR(64)                 NOT NULL,
    price       DECIMAL(6,2),
    description TEXT
);


