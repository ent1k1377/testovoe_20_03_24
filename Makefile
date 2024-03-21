.PHONY: run docker migrateinit migrateup migratedown sqlc

run:
	go run cmd/main.go

docker:
	docker compose up --build

migrateinit:
	migrate create -ext sql -dir ./db/migration -seq init_schema

migrateup:
	migrate -path db/migration -database "postgresql://root:root@localhost:5411/root?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:root@localhost:5411/root?sslmode=disable" -verbose down

sqlc:
	sqlc generate

