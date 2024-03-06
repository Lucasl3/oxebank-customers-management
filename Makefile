build:
	@go build -o bin/customers-management

# run: build
# 	@./bin/customers-management

run: build
	@CompileDaemon -command="./customers-management"

migrate:
	@go run migrate/migrate.go

test: 
	@go test -v ./...