
include .env
export
migrateup:
	migrate -path db/migration -database "mysql://root:zusammen-network@tcp(127.0.0.1:3306)/zusammen" -verbose up

migratedown:
	migrate -path db/migration -database "mysql://root:zusammen-network@tcp(127.0.0.1:3306)/zusammen" -verbose down
