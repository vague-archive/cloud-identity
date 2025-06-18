download:
	go mod download

run:
	go run cmd/server/main.go

test:
	go test ./...

build:
	CGO_ENABLED=1 go build -o bin/server ./cmd/server/main.go

clean:
	rm -rf bin

.PHONY: download run build clean
