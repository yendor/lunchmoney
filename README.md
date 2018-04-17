![Build Status](https://travis-ci.org/yendor/lunchmoney.svg?branch=master)

# Lunchmoney
A budget manager written in go, inspired by YNAB4

# Migrations
Install goose to be able to run the migrations

    go get -u github.com/pressly/goose/cmd/goose

Run the migrations with

    make migrateup 

