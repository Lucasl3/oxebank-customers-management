build:
	@go build -o bin/customers-management

run: build
	@./bin/customers-management

test: 
	@go test -v ./...