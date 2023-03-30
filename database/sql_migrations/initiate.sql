-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE users (
                        id SERIAL PRIMARY KEY,
                        username VARCHAR(256),
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
                        average_rating SMALLINT,
                        city_id BIGINT,
                        country_id BIGINT,
                        created_at TIMESTAMP,
                        updated_at TIMESTAMP
);

CREATE TABLE room (
                        id SERIAL PRIMARY KEY,
                        name VARCHAR(256),
                        description VARCHAR(256),
                        image_url VARCHAR(256),
                        price BIGINT,
                        hotel_id BIGINT,
                        created_at TIMESTAMP,
                        updated_at TIMESTAMP
);

CREATE TABLE review (
                        id SERIAL PRIMARY KEY,
                        hotel_id BIGINT NOT NULL,
                        rating SMALLINT NOT NULL,
                        description VARCHAR(256),
                        created_at TIMESTAMP,
                        updated_at TIMESTAMP
)

-- +migrate StatementEnd