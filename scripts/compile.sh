#!/bin/bash

CLIENT_SOURCE="./cmd/client/main.go"
CLIENT_TARGET="./build/client"
SERVER_SOURCE="./cmd/server/main.go"
SERVER_TARGET="./build/server"

KEY=$(openssl rand -base64 16)
VAL=$(openssl rand -base64 16)
TOKEN=$(echo ${VAL} | openssl dgst -sha256 -hmac ${KEY} -binary | openssl enc -base64 -A)
echo ${TOKEN} > ./security/token

VERSION_INITIAL="00.01.00"

# Create shortcuts for functions and static files
if [ ! -e "./functions" ] || [ ! -e "./static" ]; then
    ln -sf ./internal/http/server/functions .
    ln -sf ./internal/http/server/static .
fi 

# Initialize git
if [ ! -d ".git" ]; then
    git init
    git add .
    git commit -m "initial commit" .
    git tag ${VERSION_INITIAL}
fi

mkdir -p "./build"

FAASIFY_TOKEN=${TOKEN} go generate ./...
goimports -w ./internal/http/server/router.go
go mod tidy

# Minify and bundle static contents
rm -f ./static/bundle*
minify -r -b -o ./static/bundle.js ./static/*.js &>/dev/null
minify -r -b -o ./static/bundle.css ./static/*.css &>/dev/null

# Copy the bundle
rm -f ./bundle/*
cp -f ./static/*.htm* ./bundle/ &>/dev/null
cp -f ./static/*.json ./bundle/ &>/dev/null
cp -f ./static/*.svg ./bundle/ &>/dev/null
cp -f ./static/bundle.* ./bundle/ &>/dev/null
for FILE in $(find ./bundle/ -name "*.*"); do
    gzip -9 ${FILE}
    mv ${FILE}.gz ${FILE}
done

export CGO_ENABLED=0
FAASIFY_TOKEN=${TOKEN} go build -ldflags "\
    -s -w" \
    -o ${SERVER_TARGET} ${SERVER_SOURCE}
FAASIFY_TOKEN=${TOKEN} go build -ldflags "\
    -s -w" \
    -o ${CLIENT_TARGET} ${CLIENT_SOURCE}

# Minify binary
upx ${CLIENT_TARGET}
upx ${SERVER_TARGET}

