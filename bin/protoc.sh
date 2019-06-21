#!/bin/bash

protoc \
  -I protobuf \
  -I $GOPATH/src/github.com/googleapis/googleapis \
  --include_imports \
  --include_source_info \
  --descriptor_set_out=proxy/proto.pb \
  ./protobuf/adventar/v1/*.proto

protoc \
  -I protobuf \
  -I $GOPATH/src/github.com/googleapis/googleapis \
  --go_out plugins=grpc:grpc-server \
  ./protobuf/adventar/v1/*.proto

protoc \
  -I protobuf \
  -I $GOPATH/src/github.com/googleapis/googleapis \
  --js_out=import_style=commonjs:frontend/lib/grpc \
  --grpc-web_out=import_style=commonjs+dts,mode=grpcwebtext:frontend/lib/grpc \
  ./protobuf/adventar/v1/*.proto \
  $GOPATH/src/github.com/googleapis/googleapis/google/api/{annotations,http}.proto
