#!/bin/bash

source "/var/www/gcp/local/common.sh"

echogreen "push node-protoc to gcr"

docker build \
  --target=release \
  -t "gcr.io/$PROJECT_ID/node-protoc:latest" \
  --build-arg VERS=$PROTOC_TAG \
  .

#gcloud auth configure-docker --quiet

docker push "gcr.io/$PROJECT_ID/node-protoc:latest"
