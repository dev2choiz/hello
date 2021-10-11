#!/usr/bin/env bash

INSTANCE="$GCP_PROJECT_NAME:$GCP_REGION:$GCP_SQL_INSTANCE"

/cloud_sql_proxy \
    -instances=$INSTANCE=tcp:0.0.0.0:5432 \
    -credential_file=/credentials.json
