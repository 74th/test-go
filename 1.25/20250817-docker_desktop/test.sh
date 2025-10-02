#!/usr/local/bin/bash
# docker run -it --rm golang:1.25

cd /root
cat <<EOF > main.go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
EOF

GOOS=linux GOARCH=arm64 go build -o main_arm64 main.go
./main_arm64

GOOS=linux GOARCH=amd64 go build -o main_amd64 main.go
./main_amd64
