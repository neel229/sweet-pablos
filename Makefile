postgres:
	docker run --name postgres13 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=postgres -d postgres:13-alpine

createdb:
	docker exec -it postgres13 createdb --username=root --owner=root sweet-pablo

dropdb:
	docker exec -it postgres13 dropdb sweet-pablo

migrateup:
	 migrate -path pkg/user/db/migration -database "postgresql://root:postgres@localhost:5432/sweet-pablo?sslmode=disable" -verbose up

migratedown:
	 migrate -path pkg/user/db/migration -database "postgresql://root:postgres@localhost:5432/sweet-pablo?sslmode=disable" -verbose down

test:
	go test -race -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test