package handlers

import (
	"context"
	"github.com/dev2choiz/hello/pkg/protobuf/pingpb"
)

type PingServer struct {
	SvcName string
	pingpb.UnimplementedPingServer
}

func (p PingServer) Ping(ctx context.Context, request *pingpb.PingRequest) (*pingpb.PingResponse, error) {
	return &pingpb.PingResponse{Response: p.SvcName + " ok"}, nil
}
