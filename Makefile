DB_USER = $(USER)
DB_MAIN = restapi_script1 
DB_TEST = restapi_test1 

.PHONY: build
build:
	go build -v cmd/

.PHONY: tests
tests: build
	cd tests && go test -v

.PHONY: run
run: 
	go run cmd/main.go

db: 
	@echo "Создание базы данных $(DB_MAIN)"
	@psql -U $(DB_USER) -c "CREATE DATABASE $(DB_MAIN);"

	@echo "Создание базы данных $(DB_TEST)"
	@psql -U $(DB_USER) -c "CREATE DATABASE $(DB_TEST);"	


clean:
	rm -rf main

#.DEFAULT_GOAL := tests
.DEFAULT_GOAL := run