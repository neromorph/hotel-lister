-- +migrate Up
-- +migrate StatementBegin

ALTER TABLE hotel ALTER COLUMN phone TYPE VARCHAR(256);
ALTER TABLE room ALTER COLUMN price TYPE VARCHAR(256);

-- +migrate StatementEnd