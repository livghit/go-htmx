dev:
	@go run cmd/go-htmx/*.go

build:
	@go build -o bin/startApp cmd/go-htmx/*.go 

run:
	@./bin/startApp.exe

	
