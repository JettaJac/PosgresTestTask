#https://www.youtube.com/watch?v=LxJLuW5aUDQ

DB_USER = $(USER)
DB_MAIN = restapi_script1 
DB_TEST = restapi_test1 

.PHONY: build
build:
	go build -v cmd/main


#tests: build
#	cd tests && go test -v
cleant:
#	pwd
	sh /Users/jettajac/1_clean_1.sh
	cd /Users/jettajac/Documents/Simple_GO/PosgresTestTask
	pwd


.PHONY: run
run: 
	go run cmd/main.go

db: 
	@echo "Создание базы данных $(DB_MAIN)"
	@psql -U $(DB_USER) -c "CREATE DATABASE $(DB_MAIN);"

	@echo "Создание базы данных $(DB_TEST)"
	@psql -U $(DB_USER) -c "CREATE DATABASE $(DB_TEST);"	


.PHONY: test
test: 
	cd tests && go test
#  -v -race -timeout 30s ./ ...

clean:
	rm -rf main

#.DEFAULT_GOAL := tests
.DEFAULT_GOAL := run