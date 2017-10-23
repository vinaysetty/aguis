#!/bin/bash

mkdir -p ./public/_proto
mkdir -p ./proto/_proto

if [[ "$GOBIN" == "" ]]; then
  if [[ "$GOPATH" == "" ]]; then
    echo "Required env var GOPATH is not set; aborting with error; see the following documentation which can be invoked via the 'go help gopath' command."
    go help gopath
    exit -1
  fi

  echo "Optional env var GOBIN is not set; using default derived from GOPATH as: \"$GOPATH/bin\""
  export GOBIN="$GOPATH/bin"
fi

protoc \
  --plugin=protoc-gen-ts=./public/node_modules/.bin/protoc-gen-ts \
  --plugin=protoc-gen-go=${GOBIN}/protoc-gen-go \
  -I ./proto \
  --js_out=import_style=commonjs,binary:./public/_proto \
  --go_out=plugins=grpc:./proto/_proto \
  --ts_out=service=true:./public/_proto \
  ./proto/aguis/library/aguis_service.proto
