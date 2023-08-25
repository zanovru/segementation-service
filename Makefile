.PHONY: build
build:
	go build -o ./build/segment-service ./cmd/main.go

run: build
	./build/segment-service
