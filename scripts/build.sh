#!/bin/bash

SOURCE="./cmd/serve-http/main.go"
TARGET="./build/serve-http"

VERSION_INITIAL="00.01.00"

# Create shortcuts for functions and static files
if [ ! -e "./functions" ] || [ ! -e "./static" ]; then
    ln -sf "./internal/http/server/functions" .
    ln -sf "./internal/http/server/static" .
fi 

# Initialize git
if [ ! -d ".git" ]; then
    git init
    git add .
    git commit -m "initial commit" .
    git tag ${VERSION_INITIAL}
fi

APP=`basename $(pwd)`
BUILD=`git rev-parse --short HEAD`
VERSION=`git tag -n5 | head -1 | cut -f1 -d" "`

mkdir -p "./build"

go generate ./...
goimports -w "./internal/http/server/router.go"

go mod tidy

go build -ldflags "\
    -X 'main.app=$APP' \
    -X 'main.build=$BUILD' \
    -X 'main.version=$VERSION' \
    -s -w" \
    -o ${TARGET} ${SOURCE}

upx ${TARGET}
