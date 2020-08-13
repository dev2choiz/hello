#!/usr/bin/env bash

. ./.env

SUBSTITUTIONS="--substitutions=_CLUSTER_NAME=$CLUSTER_NAME,_CLUSTER_ZONE=$CLUSTER_ZONE,_NAMESPACE=$NAMESPACE,_APP_NAME=$APP_NAME,_IMAGE_TAG=$IMAGE_TAG,_HELM_TAG=$HELM_TAG"

echo "-------------------------------------"
echo "SUBSTITUTIONS=$SUBSTITUTIONS"
echo "-------------------------------------"
cloud-build-local --config=./deployments/cloudbuild/deployment.yaml $SUBSTITUTIONS --dryrun=false .
#gcloud builds submit --config=./deployments/cloudbuild/deployment.yaml $SUBSTITUTIONS .

