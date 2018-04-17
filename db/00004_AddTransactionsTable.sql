
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE transactions (
    id int not null,
    account_id int not null,
    category_id int not null,
    occurred datetime,
    payee text,
    memo text
    debit int
    credit int
    is_cleared tinyint(1) not null default 0,
    is_reconciled tinyint(1) not null default 0,
    PRIMARY KEY (id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS transactions;
