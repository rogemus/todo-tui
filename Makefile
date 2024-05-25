.PHONY: build run clean

build:
	go build -o todotui ./cmd/tui/main.go

run: build
	./todotui

clean:
	go clean
	rm -f todotui

