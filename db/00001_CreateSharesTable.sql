-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE shares (
    id integer NOT NULL,
    stockcode character varying(10) NOT NULL,
    qty integer NOT NULL,
    PRIMARY KEY (id)
);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
