#!/bin/bash

if [ $# -ne 1 ]; then
    echo "Usage: $0 <new_directory_name>"
    exit 1
fi

mkdir $1
cd $1
go mod init github.com/74th/go-testing/$1
go work use $1