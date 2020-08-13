#!/usr/bin/env bash

. ./.env

SUBSTITUTIONS="--substitutions=_PROJECT_ID=$PROJECT_ID,_APP_NAME=$APP_NAME,_IMAGE_TAG=$IMAGE_TAG"
echo "-------------------------------------"
echo "SUBSTITUTIONS=$SUBSTITUTIONS"
echo "-------------------------------------"
cloud-build-local --config=./deployments/cloudbuild/build-image.yaml "$SUBSTITUTIONS" --dryrun=false .
#gcloud builds submit --config=./deployments/cloudbuild/build-image.yaml "$SUBSTITUTIONS" .
