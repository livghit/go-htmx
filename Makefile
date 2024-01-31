# You can use this or air 
dev:
	@go run cmd/go-htmx/*.go

build:
	@go build -o bin/main cmd/go-htmx/*.go 

run:
	@./bin/startApp.exe
