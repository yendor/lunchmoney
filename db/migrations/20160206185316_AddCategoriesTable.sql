
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE categories (
    id int NOT NULL,
    name text,
    PRIMARY KEY(id)
);



-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS categories;


