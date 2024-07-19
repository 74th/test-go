#!/bin/bash
set -xe

if [ $# -ne 1 ]; then
    echo "Usage: $0 <new_directory_name>"
    exit 1
fi

TODAY=$(date +%Y%m%d)
DIR_NAME="${TODAY}_$1"

mkdir -p $DIR_NAME
cp template/main_test.go ./$DIR_NAME/
cd $DIR_NAME
go mod init github.com/74th/testing-go/$DIR_NAME
cd ..
go work use $DIR_NAME