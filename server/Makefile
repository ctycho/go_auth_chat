.postgresinit:
	docker run --name postgres16 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:16-alpine

.postgres:
	docker exec -it postgres16 psql

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root go-chat 

dropdb:
	docker exec -it postgres16 dropdb go-chat 

migrate:
	migrate create -ext sql -dir db/migrations add_users_table

migrateup:
	migrate -path db/migrations -database "postgres://root:password@localhost:5433/go-chat?sslmode=disable" -verbose up

migrateup:
	migrate -path db/migrations -database "postgres://root:password@localhost:5433/go-chat?sslmode=disable" -verbose down

.PHONY: postgresinit postgres createdb dropdb migrate migrateup migrateup
