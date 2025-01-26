all: build
.PHONY: build run clean

build:
	go build -o bin/main cmd/*.go

run:
	go run cmd/main.go

clean:
	rm -rf bin

gen_proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/quill.proto