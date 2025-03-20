# Banking Application Excercise
This application is made as an excercise to perform simple banking transaction. any feature on this app already handle race condition issue on the transaction (via database lock)

## Requirement
to be able to run this application locally, you will need the following:
- go version 1.23
- open api code generator ([Installation Docs](https://github.com/oapi-codegen/oapi-codegen))
- mockery
- docker with docker compose installed ([Installation Docs](https://docs.docker.com/engine/install/))
- GNU make ([Docs](https://www.gnu.org/software/make/))

## Local Development
```
$ go mod download
...
$ go run ./cmd/api/...
```

## Available Make Command
- `make generate`: generate openapi code generator for echo server and request/response datatypes
- `make generate_mock`: generate mock code using mockery
- `make test_unit`: run unit test (without integration test)
- `make test_integration`: run full test (make sure local db is running)