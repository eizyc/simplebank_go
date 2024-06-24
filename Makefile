DB_URL=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable

postgres:
		docker run --name postgres16 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

createdb:
		docker exec -it postgres16 createdb --username=root --owner=root simple_bank

dropdb:
		docker exec -it postgres16 dropdb simple_bank

migrateup:
		migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
		migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

sqlc: 
		sqlc generate

test:
		go test -v -cover ./...

server:
		go run main.go

mock:
		mockgen -package mockdb -destination db/mock/store.go github.com/eizyc/simplebank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 new_migration sqlc test server mock