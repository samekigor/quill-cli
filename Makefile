all: build
.PHONY: build run clean

gen_proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/auth/auth.proto

build:
	go build -o quill main.go

install:
	go build -o quill main.go
	mkdir -p $(GOBIN)
	mv quill $(GOBIN)/quill
