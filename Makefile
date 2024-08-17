createdb:
	docker exec -it bank-postgres createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it bank-postgres dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:password@localhost:5433/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:password@localhost:5433/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/hoangduc02011998/simple_bank/db/sqlc Store

.PHONY: createdb dropdb migrateup migratedown server test sqlc mock