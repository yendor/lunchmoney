-- +goose Up
CREATE TABLE accounts (
    id int NOT NULL,
    name text,
    currency text,
    currency_symbol_left text,
    currency_symbol_right text,
    icon text,
    is_active tinyint(1) not null default 0,
    cleared_total int,
    total int,
    PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE IF EXISTS accounts;