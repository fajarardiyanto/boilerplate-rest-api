tidy:
	@go mod tidy
test:
	@go test -v ./...
run:
	@go run ./cmd .