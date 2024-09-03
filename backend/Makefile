MAIN_PATH=cmd/main.go
BIN_PATH=dist/app

test:
	go test -v ./...

build: test
	go build -o $(BIN_PATH) $(MAIN_PATH)

start:
	./$(BIN_PATH)

run:
	go run $(MAIN_PATH)

migrate:
	go run $(MAIN_PATH) -migrate

help:
	go run $(MAIN_PATH) -help
