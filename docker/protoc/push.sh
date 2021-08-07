#!/bin/bash

source "/var/www/gcp/local/common.sh"

echogreen "push protoc to gcr"

docker build \
  -t "gcr.io/$PROJECT_ID/protoc:latest" \
  --target=release \
  --build-arg VERS=$PROTOC_TAG \
  .

#gcloud auth configure-docker --quiet

docker push "gcr.io/$PROJECT_ID/protoc:latest"
