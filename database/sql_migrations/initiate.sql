-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE user (
                        id SERIAL PRIMARY KEY,
                        name VARCHAR(256),
                        email VARCHAR(256),
                        password VARCHAR(256),
                        created_at TIMESTAMP,
                        updated_at TIMESTAMP
);

CREATE TABLE country (
                        id BIGINT NOT NULL,
                        name VARCHAR(256),
                        created_at TIMESTAMP,
                        updated_at TIMESTAMP
);

CREATE TABLE city (
                        id BIGINT NOT NULL,
                        name VARCHAR(256),
                        country_id BIGINT,
                        created_at TIMESTAMP,
                        updated_at TIMESTAMP
);

CREATE TABLE hotel (
                        id SERIAL PRIMARY KEY,
                        name VARCHAR(256),
                        description VARCHAR(256),
                        image_url VARCHAR(256),
                        address VARCHAR(256),
                        phone BIGINT,
                        email VARCHAR(256),
                        website VARCHAR(256),
                        rating BIGINT,
                        created_at TIMESTAMP,
                        updated_at TIMESTAMP,
                        city_id BIGINT
);

CREATE TABLE room (
                        id SERIAL PRIMARY KEY,
                        name VARCHAR(256),
                        description VARCHAR(256),
                        image_url VARCHAR(256),
                        price BIGINT,
                        created_at TIMESTAMP,
                        updated_at TIMESTAMP,
                        hotel_id BIGINT
);

-- +migrate StatementEnd