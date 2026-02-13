BIN = bin/gor

build:
	go build -o $(BIN) ./cmd/gor/

run: build
	./$(BIN)

test:
	go test ./test/ -count=1

test-v:
	go test ./test/ -count=1 -v

benchmark: build
	@bash benchmark/run.sh

fmt:
	go fmt ./...

clean:
	rm -f $(BIN)

.PHONY: build run test test-v benchmark fmt clean

