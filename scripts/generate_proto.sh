#!/usr/bin/env bash

echo "Generate protobuf"
rm -rf pkg/protobuf/*

#protoc --proto_path=api/proto \
#  --proto_path=third_party \
#  --go_out=. \
#  --go-grpc_out=. \
#  --grpc-gateway_out=. \
#  --grpc-gateway_opt logtostderr=true \
#  --grpc-gateway_opt grpc_api_configuration=api/v1/proto/gateway.yaml \
#  --openapiv2_out=api/v1/openapi \
#  --openapiv2_opt logtostderr=true \
#  --openapiv2_opt generate_unbound_methods=true \
#  --openapiv2_opt grpc_api_configuration=api/v1/proto/gateway.yaml \
#  check-svc.proto

protoc --proto_path=api/proto \
  --proto_path=third_party \
  --go_out=. \
  --go-grpc_out=. \
  --grpc-gateway_out=. \
  health.proto notify.proto ping.proto
