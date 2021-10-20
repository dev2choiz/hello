#!/usr/bin/env bash

OUT_DIR=./src/protobuf
IMPORT_STYLE="commonjs"

echo "Generate protobuf in $OUT_DIR"

rm -rf $OUT_DIR
mkdir -p $OUT_DIR

protoc --proto_path=./../api/proto \
  --proto_path=./../third_party \
  --js_out="import_style=${IMPORT_STYLE},binary:${OUT_DIR}/" \
  --ts_out="service=grpc-web:${OUT_DIR}/" \
  health.proto notify.proto ping.proto sandbox.proto \
  google/api/annotations.proto google/api/http.proto
