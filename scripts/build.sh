#!/bin/bash

APP=`basename $(pwd)`
BUILD=`git rev-parse --short HEAD`
REGISTRY="greenfield.azurecr.io"
VERSION=`git tag -n5 | head -1 | cut -f1 -d" "`

# Build podman image
podman build \
	-f ./build/Dockerfile \
	-t ${APP}:${VERSION} \
	.

podman image tag ${APP}:${VERSION} ${REGISTRY}/${APP}:${VERSION}
podman push ${REGISTRY}/${APP}:${VERSION}

