#!/bin/bash

set -e

docker build -t adventar-envoy .
docker run --rm -it -p 8000:8000 adventar-envoy /usr/local/bin/envoy -c /etc/envoy/envoy-local.yaml
