#!/usr/bin/env bash

set -xe

mkdir -p out

cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" out/
GOOS=js GOARCH=wasm go build -o out/main.wasm
