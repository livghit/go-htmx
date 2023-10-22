dev:
	@go run cmd/linkhub/*.go

build:
	@go build -o bin/startApp cmd/main.go 

run:
	@./bin/startApp.exe

	
