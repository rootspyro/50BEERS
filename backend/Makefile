MAIN_PATH=cmd/main.go
BIN_PATH=dist/app

build:
	go build -o $(BIN_PATH) $(MAIN_PATH)

start:
	./$(BIN_PATH)

run:
	go run $(MAIN_PATH)

migrate:
	go run $(MAIN_PATH) -migrate
