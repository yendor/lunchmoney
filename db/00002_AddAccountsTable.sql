-- +goose Up
CREATE TABLE accounts (
    id INTEGER PRIMARY KEY,
    name text,
    currency_code text(3),
    currency_symbol_left text not null default '',
    currency_symbol_right text not null default '',
    decimal_places int not null default 2,
    icon text,
    is_active tinyint(1) not null default 0,
    cleared_total int not null default 0,
    total int int not null default 0
);

-- +goose Down
DROP TABLE IF EXISTS accounts;