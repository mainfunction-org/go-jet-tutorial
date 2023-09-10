CREATE USER jet;

CREATE DATABASE pet_store;

GRANT ALL PRIVILEGES ON DATABASE pet_store TO jet;

\c pet_store;

CREATE TYPE pet_status AS ENUM ('available', 'pending', 'sold');

CREATE TABLE pet
(
    id     BIGSERIAL PRIMARY KEY,
    name   TEXT       NOT NULL,
    status pet_status NOT NULL DEFAULT 'available'::pet_status
);

CREATE TABLE category
(
    id   BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE pet_category
(
    pet_id BIGINT NOT NULL REFERENCES pet(id),
    category_id BIGINT NOT NULL REFERENCES category(id)
);

CREATE TABLE tag
(
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE pet_tag
(
    pet_id BIGINT NOT NULL REFERENCES pet(id),
    tag_id BIGINT NOT NULL REFERENCES tag(id)
);


