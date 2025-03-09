.PHONY: all build run clean

build:
	go build -o run ./cmd/main.go

run: build
	./run

clean:
	rm -f run