CREATE TABLE orders(
    order_uid VARCHAR UNIQUE NOT NULL,
    data JSON NOT NULL
);