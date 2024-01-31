# Mailer Micro Service Makefile

# lint: runs a golangci-lint with the same settings as in the CI.
lint:
	golangci-lint run ./...

# check: executes a static check.
check:
	staticcheck ./...

# test: executes a test suite.
test:
	go test ./...

# build: builds application.
build:
	env GOOS=linux GOARCH=amd64 go build -o mailer-micro-service main.go

# run: starts an app.
run:
	go run main.go
