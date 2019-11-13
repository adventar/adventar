#!/bin/bash

cd $(dirname $0)

# envoy
protoc \
  --include_imports \
  --include_source_info \
  --descriptor_set_out=../api-server/envoy/proto.pb \
  adventar/v1/*.proto

# grpc-server
protoc \
  --go_out plugins=grpc:../api-server/grpc-server/grpc \
  adventar/v1/*.proto

# frontend
protoc \
  --js_out=import_style=commonjs:../frontend/lib/grpc \
  --grpc-web_out=import_style=commonjs+dts,mode=grpcwebtext:../frontend/lib/grpc \
  adventar/v1/*.proto google/api/*.proto

# hot fix
gsed -i -e "s/extend(proto/extend(exports/" ../frontend/lib/grpc/adventar/v1/adventar_pb.js
