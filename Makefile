build:
	@go build -o bin/distributed

run:
	@./bin/distributed

test:
	@go test ./... -v