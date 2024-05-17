#use this or air to tun and build
dev:
	@go run cmd/go-htmx/*.go

build:
	go build -o bin/main cmd/go-htmx/*.go  && bin/main

seed:
	@go run cmd/seed/*.go
