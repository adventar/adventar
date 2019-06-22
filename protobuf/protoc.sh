#!/bin/bash

cd $(dirname $0)

# proxy
protoc \
  --include_imports \
  --include_source_info \
  --descriptor_set_out=../proxy/proto.pb \
  adventar/v1/*.proto

# grpc-server
protoc \
  --go_out plugins=grpc:../grpc-server \
  adventar/v1/*.proto

# frontend
protoc \
  --js_out=import_style=commonjs:../frontend/lib/grpc \
  --grpc-web_out=import_style=commonjs+dts,mode=grpcwebtext:../frontend/lib/grpc \
  adventar/v1/*.proto google/api/*.proto
