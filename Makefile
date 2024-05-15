DB_USER = $(USER)
DB_MAIN = restapi_script
DB_TEST = restapi_test



build:
	go build -v cmd/

tests: build
	cd tests && go test -v

run: 
	go run cmd/main.go

db: 
	@echo "Создание базы данных $(DB_MAIN)"
	@psql -U $(DB_USER) -c "CREATE DATABASE $(DB_MAIN);"

	@echo "Создание базы данных $(DB_TEST)"
	@psql -U $(DB_USER) -c "CREATE DATABASE $(DB_TEST);"	

migrate:
	migrate -path ./migrations -database 'postgres://user:password@0.0.0.0:5436/restapi_script?sslmode=disable' up
	migrate -path ./migrations -database 'postgres://user:password@0.0.0.0:5436/restapi_test?sslmode=disable' up

migrateup:
	migrate -path migrations -database "postgres://localhost/restapi_script?sslmode=disable" up
	migrate -path migrations -database "postgres://localhost/restapi_test?sslmode=disable" up

migratedrop:	
	migrate -path migrations -database "postgres://localhost/restapi_script?sslmode=disable" drop
	migrate -path migrations -database "postgres://localhost/restapi_test?sslmode=disable" drop


clean:
	rm -rf main

#.DEFAULT_GOAL := tests
.DEFAULT_GOAL := run
.PHONY: build tests run tests db migrate clean