createdb:
	createdb --username=postgres --owner=postgres go_finance

postgres:
	docker run -e POSTGRES_USER=devmarcos -e POSTGRES_PASSWORD=karybe021313 -e POSTGRES_DB=go-finance -e POSTGRES_HOST=go-finance.cayl5ldqzc4i.us-east-1.rds.amazonaws.com \-p 5432:5432 \--name gofinace \-d postgres:14-alpine
migrateup:
	migrate -path db/migration -database "postgresql://devmarcos:karybe021313@go-finance.cayl5ldqzc4i.us-east-1.rds.amazonaws.com:5432/postgres?sslmode=require" -verbose up

migrationdrop:
	migrate -path db/migration -database "postgresql://devmarcos:karybe021313@go-finance.cayl5ldqzc4i.us-east-1.rds.amazonaws.com:5432/postgres?sslmode=require" -verbose down

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: createdb postgres dropdb migrateup migrationdrop test server