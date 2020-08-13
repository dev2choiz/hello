#!/bin/bash

. ./.env

gcloud builds submit . --config=./cloud/builds/unit-test.yaml
