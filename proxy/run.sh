#!/bin/bash

set -e

protoc \
  -I ../protobuf \
  -I $GOPATH/src/github.com/googleapis/googleapis \
  --include_imports \
  --include_source_info \
  --descriptor_set_out=proto.pb \
  ../protobuf/adventar/v1/*.proto
docker build -t adventar-grpc-web-proxy .
docker run --rm -it -p 8000:8000 adventar-grpc-web-proxy
