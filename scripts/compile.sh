#!/bin/bash

SERVER_SOURCE="./cmd/server/main.go"
SERVER_TARGET="./build/server"

KEY=$(openssl rand -base64 16)
VAL=$(openssl rand -base64 16)
TOKEN=$(echo ${VAL} | openssl dgst -sha256 -hmac ${KEY} -binary | openssl enc -base64 -A)
echo ${TOKEN} > ./security/token

# Create shortcuts for functions and static files
if [ ! -e "./functions" ] || [ ! -e "./static" ]; then
    ln -sf ./internal/http/server/functions .
    ln -sf ./internal/http/server/static .
fi 

# Rename the module if needed
MODULE_OLD=$(cat go.mod | head -1 | cut -f2 -d' ')
BASENAME_OLD=$(basename ${MODULE_OLD})
BASENAME_NEW=$(basename ${PWD})
if [ ! "${BASENAME_OLD}" == "${BASENAME_NEW}" ]; then
    rm -rf go.* && go mod init &>/dev/null
    MODULE_NEW=$(cat go.mod | head -1 | cut -f2 -d' ')
    for FILE in $(find . -name "*.go"); do
        sed -i 's|'${MODULE_OLD}'|'${MODULE_NEW}'|g' $FILE
    done
fi

mkdir -p "./build"
rm -rf "./vendor"
FAASIFY_TOKEN=${TOKEN} go generate ./...
goimports -w ./internal/http/server/router.go
go install github.com/tdewolff/minify/v2/cmd/minify@latest
go mod tidy
go mod vendor

# Minify and bundle static contents
rm -f ./static/bundle*
minify -r -b -o ./static/bundle.js ./static/*.js &>/dev/null
minify -r -b -o ./static/bundle.css ./static/*.css &>/dev/null

# Copy the bundle
rm -f ./bundle/*
cp -f ./static/*.htm* ./bundle/ &>/dev/null
cp -f ./static/*.ico ./bundle/ &>/dev/null
cp -f ./static/*.json ./bundle/ &>/dev/null
cp -f ./static/*.png ./bundle/ &>/dev/null
cp -f ./static/*.svg ./bundle/ &>/dev/null
cp -f ./static/bundle.* ./bundle/ &>/dev/null
for FILE in $(find ./bundle/ -name "*.*"); do
    gzip -9 ${FILE}
    mv ${FILE}.gz ${FILE}
done

for ARCH in arm64 amd64; do
    for OS in darwin linux; do
        CGO_ENABLED=0 GOARCH=${ARCH} GOOS=${OS} FAASIFY_TOKEN=${TOKEN} go build -ldflags "\
            -s -w" \
            -o ${SERVER_TARGET}_${OS}_${ARCH} ${SERVER_SOURCE}
        # Minify binary
        upx ${SERVER_TARGET}_${OS}_${ARCH}
    done
done
