#!/usr/bin/env bash

echo "Generate protobuf"
rm -rf pkg/protobuf/*

protoc --proto_path=api/proto \
  --proto_path=third_party \
  --include_imports \
  --go_out=. \
  --go-grpc_out=. \
  --descriptor_set_out=./api/config/api_descriptor.pb \
  health.proto notify.proto sandbox.proto

protoc --proto_path=api/proto \
  --proto_path=third_party \
  --go_out=. \
  --go-grpc_out=. \
  ping.proto
