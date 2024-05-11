#https://www.youtube.com/watch?v=LxJLuW5aUDQ

DB_USER = jettajac
DB_NAME = your_database_name  

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

bd: 
	@echo "Создание базы данных $(DB_NAME)"
	@psql -U $(DB_USER) -c "CREATE DATABASE $(DB_NAME);"


.PHONY: test
test: 
# пока запускаем с сервера, возможно перенести в другую папку
	cd internal/server && go test
# -v -race -timeout 30s ./ ...

clean:
	rm -rf main

#.DEFAULT_GOAL := tests
.DEFAULT_GOAL := run