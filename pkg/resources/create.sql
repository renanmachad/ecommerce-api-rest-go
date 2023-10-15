create schema if not exists goland;


CREATE TABLE products
(
    id          INT            NOT NULL generated always as identity,
    name        VARCHAR(255)   NOT NULL,
    description TEXT,
    price       DECIMAL(10, 2) NOT NULL,
    image       VARCHAR(255),
    PRIMARY KEY (id)
);

CREATE TABLE categories
(
    id   INT          NOT NULL generated always as identity,
    name VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE  if not exists product_categories
(
    product_id  INT NOT NULL,
    category_id INT NOT NULL,
    PRIMARY KEY (product_id, category_id),
    FOREIGN KEY (product_id) REFERENCES products (id),
    FOREIGN KEY (category_id) REFERENCES categories (id)
);

CREATE TABLE if not exists customers
(
    id       INT          NOT NULL generated always as identity,
    name     VARCHAR(255) NOT NULL,
    email    VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE if not exists orders
(
    id           INT            NOT NULL generated always as identity,
    customer_id  INT            NOT NULL,
    total_amount DECIMAL(10, 2) NOT NULL,
    order_status VARCHAR(255)   NOT NULL,
    created_at   TIMESTAMP      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (customer_id) REFERENCES customers (id)
);

CREATE TABLE if not exists order_items
(
    order_id   INT NOT NULL,
    product_id INT NOT NULL,
    quantity   INT NOT NULL,
    PRIMARY KEY (order_id, product_id),
    FOREIGN KEY (order_id) REFERENCES orders (id),
    FOREIGN KEY (product_id) REFERENCES products (id)
);