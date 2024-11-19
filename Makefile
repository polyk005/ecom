build:
	@go build -o bin/ecom cmd/main.go

test:
	@test -v ./...

run: build
	@./bin/ecom