![Build Status](https://travis-ci.org/yendor/lunchmoney.svg?branch=master)

# Lunchmoney
A budget manager written in go, inspired by YNAB4

# Migrations
Copy the example config into place

    cp example.config.json config.json

Install goose to be able to run the migrations

    go get -u github.com/pressly/goose/cmd/goose

Run the migrations with

    make migrateup 

Build the program

    make

Run the program 

    ./lunchmoney