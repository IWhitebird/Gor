.PHONY: build run test test-v fmt clean

BIN = bin/gor

build:
	go build -o $(BIN) ./cmd/gor/

run: build
	./$(BIN)

test:
	go test ./test/ -count=1 -v

fmt:
	go fmt ./...

clean:
	rm -f $(BIN)

