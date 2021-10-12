#!/usr/bin/env bash

echo "Generate protobuf"

OUT_DIR=./protobuf

rm -rf $OUT_DIR
mkdir -p $OUT_DIR

protoc --proto_path=./../api/proto \
  --proto_path=./../third_party \
  --js_out="import_style=commonjs,binary:${OUT_DIR}/" \
  --ts_out="service=grpc-web:${OUT_DIR}/" \
  health.proto notify.proto ping.proto sandbox.proto \
  google/api/annotations.proto google/api/http.proto
