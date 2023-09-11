CREATE
USER jet;

CREATE DATABASE pet_store;

\c pet_store;

CREATE TYPE pet_status AS ENUM ('available', 'pending', 'sold');

CREATE TABLE pets
(
    id     BIGSERIAL PRIMARY KEY,
    name   TEXT       NOT NULL,
    status pet_status NOT NULL DEFAULT 'available'::pet_status
);

CREATE TABLE categories
(
    id   BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE pet_categories
(
    pet_id      BIGINT NOT NULL REFERENCES pets (id),
    category_id BIGINT NOT NULL REFERENCES categories (id)
);

CREATE TABLE tags
(
    id   BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE pet_tags
(
    pet_id BIGINT NOT NULL REFERENCES pets (id),
    tag_id BIGINT NOT NULL REFERENCES tags (id)
);

CREATE TYPE user_status AS ENUM ('active', 'disabled');

CREATE TABLE users
(
    id         BIGSERIAL PRIMARY KEY,
    username   VARCHAR(20)  NOT NULL,
    first_name VARCHAR(100),
    last_name  VARCHAR(100),
    email      VARCHAR(100) NOT NULL,
    password   VARCHAR(32)  NOT NULL,
    phone      VARCHAR(20),
    status     user_status DEFAULT 'active'::user_status
);

CREATE TYPE order_status AS ENUM ('placed', 'approved', 'delivered');

CREATE TABLE orders
(
    id BIGSERIAL PRIMARY KEY,
    pet_id BIGINT NOT NULL REFERENCES pets (id),
    quantity INT NOT NULL DEFAULT 1,
    ship_date TIMESTAMP,
    status order_status DEFAULT 'placed'::order_status,
    complete BOOLEAN DEFAULT false,
    created_at TIMESTAMP DEFAULT NOW()
);
