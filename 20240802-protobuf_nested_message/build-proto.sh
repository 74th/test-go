#!/bin/bash
set -xe

mkdir -p bin/
mkdir -p pb/

if [ ! -f "bin/protoc-gen-go-grpc" ]; then
    export GOBIN=$(pwd)/bin
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
fi
export PATH=$(pwd)/bin:$PATH

protoc --go_out=pb/ --go_opt=paths=source_relative --go-grpc_out=pb/ --go-grpc_opt=paths=source_relative *.proto