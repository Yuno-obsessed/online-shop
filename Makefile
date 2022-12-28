
include .env
export
migrateup:
	migrate -path db/migration -database "mysql://root:library-watcher@tcp(127.0.0.1:3306)/library" -verbose up

migratedown:
	migrate -path db/migration -database "mysql://root:library-watcher@tcp(127.0.0.1:3306)/library" -verbose down
