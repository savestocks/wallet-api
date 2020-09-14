# wallet-api

## Golang Spell
The project was initialized using [Golang Spell](https://github.com/golangspell/golangspell).

## Architectural Model
The Architectural Model adopted to structure the application is based on The Clean Architecture.
Further details can be found here: [The Clean Architecture](https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html) and in the Clean Architecture Book.

## Environment variables

```sh
export PORT=8080

export APP_NAME=wallet-api

export LOG_LEVEL=INFO
```

## Dependency Management
The project is using [Go Modules](https://blog.golang.org/using-go-modules) for dependency management
Module: github.com/andersonlira/wallet-api

## Test and coverage

Run the tests

```sh 
TESTRUN=true go test ./... -coverprofile=cover.out

go tool cover -html=cover.out
```

Install [golangci-lint](https://github.com/golangci/golangci-lint#install) and run lint:

```sh
golangci-lint run
```

## Docker Build

```sh
docker build .
```