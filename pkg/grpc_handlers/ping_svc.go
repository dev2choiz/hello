package grpc_handlers

import (
	"context"
	"github.com/dev2choiz/hello/pkg/protobuf/pingpb"
)

type PingSvcServer struct {
	SvcName string
	pingpb.UnimplementedPingSvcServer
}

func (p PingSvcServer) Ping(ctx context.Context, request *pingpb.PingRequest) (*pingpb.PingResponse, error) {
	return &pingpb.PingResponse{Response: p.SvcName + " ok" }, nil
}
