run: build
	./bin/tcm-backend

build:
	go build -o bin/tcm-backend cmd/api/main.go

up:
	go run cmd/api/main.go