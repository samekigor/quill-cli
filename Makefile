all: build
.PHONY: build run clean

build:
	go build -o quill

run:
	go run main.go

clean:
	rm -rf bin

gen_proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/quill.proto

install:
    go install -o ~/go/bin/quill ./cmd/*.go
