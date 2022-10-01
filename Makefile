tidy:
	@go mod tidy
test:
	@go test -v ./...
make run:
	@go run ./cmd .