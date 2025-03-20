generate:
	go get github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen
	go generate ./api/...
	go mod tidy

generate_mock:
	mockery

test_unit:
	go test -short ./...

test_integration:
	go test ./...
