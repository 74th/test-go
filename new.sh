#!/bin/bash
set -xe

if [ $# -ne 1 ]; then
    echo "Usage: $0 <new_directory_name>"
    exit 1
fi

mkdir -p $1
cp template/main_test.go ./$1/
cd $1
go mod init github.com/74th/test-go/$1
cd ..
go work use $1