migratedb:
	migrate	-database "postgres://postgres:1234@localhost:5432/simplebank?sslmode=disable" -path db/migration up
migratedbdown:
	migrate	-database "postgres://postgres:1234@localhost:5432/simplebank?sslmode=disable" -path db/migration down

sqlc:
	docker run --rm -v "${CURDIR}:/src" -w /src kjconroy/sqlc generate

test:
	go test -v -cover ./...


.PHONY: migratedb sqlc test