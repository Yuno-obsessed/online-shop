include .env
export
BINARY_PATH=build/zusammen

build:
	mkdir -p build/
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_PATH}-darwin main.go
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_PATH}-linux main.go
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_PATH}-windows main.go

run:
	./${BINARY_PATH}

build_and_run:
	build run

clean:
	go clean
	rm ${BINARY_PATH}-darwin
	rm ${BINARY_PATH}-linux
	rm ${BINARY_PATH}-windows

zusammen-work:
	@ docker-compose up -d


	@ until mysql --host=$(MYSQL_HOST) --port=$(MYSQL_PORT) --user=$(MYSQL_USER) -p$(MYSQL_PASSWORD) --protocol=tcp -e 'SELECT 1' >/dev/null 2>&1 && exit 0; do \
	  >&2 echo "MySQL is unavailable - sleeping"; \
	  sleep 5 ; \
	done


	@ echo "MySQL is up and running!"


migrate-setup:
	@if [ -z "$$(which migrate)" ]; then echo "Installing migrate command..."; go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate; fi


migrate-up: migrate-setup
	@ migrate -database 'mysql://$(MYSQL_DSN)?multiStatements=true' -path migrations up $(N)


migrate-down: migrate-setup
	@ migrate -database 'mysql://$(MYSQL_DSN)?multiStatements=true' -path migrations down $(N)


migrate-to-version: migrate-setup
	@ migrate -database 'mysql://$(MYSQL_DSN)?multiStatements=true' -path migrations goto $(V)


drop-db: migrate-setup
	@ migrate -database 'mysql://$(MYSQL_DSN)?multiStatements=true' -path migrations drop


force-version: migrate-setup
	@ migrate -database 'mysql://$(MYSQL_DSN)?multiStatements=true' -path migrations force $(V)


migration-version: migrate-setup
	@ migrate -database 'mysql://$(MYSQL_DSN)?multiStatements=true' -path migrations version


migrateup:
	migrate -path db/migration -database "mysql://root:zusammen-network@tcp(127.0.0.1:3306)/zusammen" -verbose up

migratedown:
	migrate -path db/migration -database "mysql://root:zusammen-network@tcp(127.0.0.1:3306)/zusammen" -verbose down
