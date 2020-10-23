#!/bin/bash
#protoc --go-grpc_out=plugins=protoc-gen-go:. hello.proto

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative hello.proto
