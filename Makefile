#https://www.youtube.com/watch?v=LxJLuW5aUDQ
.PHONY: build
build:
	go build -v cmd/main

#tests: build
#	cd tests && go test -v

start: 
	go run cmd/main.go

.PHONY: test
test: 
	cd internal/app && go test
# -v -race -timeout 30s ./ ...

clean:
	rm -rf main

#.DEFAULT_GOAL := tests
.DEFAULT_GOAL := start