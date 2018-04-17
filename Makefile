.PHONY: migrateup clean

lunchmoney: *.go
	go build

clean: 
	rm lunchmoney

migrateup:
	cd db; ${GOPATH}/bin/goose sqlite3 ../lunchmoney.db up
