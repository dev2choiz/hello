#!/bin/bash

source "/var/www/gcp/local/common.sh"

echo "push helm to gcr"

echogreen "clone https://github.com/GoogleCloudPlatform/cloud-builders-community.git"
git clone https://github.com/GoogleCloudPlatform/cloud-builders-community.git
cd cloud-builders-community/helm
echogreen "build helm image + tag as gcr.io/$PROJECT_ID/helm:$HELM_TAG"
docker build \
    -t "gcr.io/$PROJECT_ID/helm:$HELM_TAG" \
    --build-arg HELM_VERSION=v$HELM_TAG \
    .


echogreen "docker push gcr.io/$PROJECT_ID/helm:$HELM_TAG"

#gcloud auth configure-docker --quiet

docker push "gcr.io/$PROJECT_ID/helm:$HELM_TAG"
cd ../..
rm -rf cloud-builders-community
