#!/usr/bin/env bash

REPORT_DIR="./reports"

mkdir -p "$REPORT_DIR"
ARGS_TEST="--coverprofile=$REPORT_DIR/cover.out"

go test "$ARGS_TEST" ./...

go tool cover -html=$REPORT_DIR/cover.out -o $REPORT_DIR/cover.html

gocover-cobertura < $REPORT_DIR/cover.out > $REPORT_DIR/cover.xml
