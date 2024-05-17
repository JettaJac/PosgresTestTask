DB_USER = $(USER)
DB_MAIN = restapi_script
DB_TEST = restapi_test
DATABASE_HOST=localhost

ifeq ($(env),local)
    DATABASE_HOST=localhost
endif
ifeq ($(env),dev)
    DATABASE_HOST=db
endif
ifeq ($(env),prod)
    DATABASE_HOST=prod-host
endif
	

build:
	chmod 777 init-scripts/create_db.sh
	cd ..

	DATABASE_HOST=${DATABASE_HOST} go build -v cmd/main.go

tests: build
	cd tests && go test -v

testsall: tests	
		cd internal/storage/sqlstore && go test
		cd internal/storage/teststore && go test

run: 
	# echo $(env) 
	DATABASE_HOST=${DATABASE_HOST} go run cmd/main.go

db: 
	@echo "Создание базы данных $(DB_MAIN)"
	@psql -U $(DB_USER) -c "CREATE DATABASE $(DB_MAIN);"

	@echo "Создание базы данных $(DB_TEST)"
	@psql -U $(DB_USER) -c "CREATE DATABASE $(DB_TEST);"	

migrate:
	migrate -path ./migrations -database 'postgres://user:password@0.0.0.0:5432/restapi_script?sslmode=disable' up
	migrate -path ./migrations -database 'postgres://user:password@0.0.0.0:5432/restapi_test?sslmode=disable' up

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
.PHONY: build tests run tests testall db migrate migrateup migratedrop clean