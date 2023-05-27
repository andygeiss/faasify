#!/bin/bash

PREFIX=`pwd | cut -f4- -d"/"`

podman run \
    -v $HOME/go:/go:Z,U \
    tinygo/tinygo:latest \
    tinygo build \
        -target=wasi \
        -o "/${PREFIX}/functions/wasm_demo/module/fn.wasm" \
        "/${PREFIX}/functions/wasm_demo/module/fn.go"

sudo chown \
    -R $USER:$USER \
    $HOME/go

