#!/bin/bash

set -e

docker build -t adventar-grpc-web-proxy .
docker run --rm -it -p 8000:8000 adventar-grpc-web-proxy /usr/local/bin/envoy -c /etc/envoy/envoy-local.yaml
