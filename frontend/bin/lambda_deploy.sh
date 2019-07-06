#!/bin/bash

yarn install
BUILD_MODE=universal API_BASE_URL=https://grpc-dev.adventar.org yarn run build
tsc --outDir .nuxt nuxt.config.ts
rm -rf handler.zip
zip handler.zip -rq lambda.js .nuxt node_modules
envchain aws-personal aws s3 cp handler.zip s3://hokaccha/lambda/adventarDevSsr/handler.zip
envchain aws-personal aws lambda update-function-code \
  --function-name adventarDevSsr \
  --s3-bucket hokaccha \
  --s3-key lambda/adventarDevSsr/handler.zip
