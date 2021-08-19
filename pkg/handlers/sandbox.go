package handlers

import (
	"context"
	"fmt"
	"github.com/dev2choiz/hello/pkg/protobuf/sandboxpb"
	"log"
)

type SandboxServer struct {
	sandboxpb.UnimplementedSandboxServer
}

func (s SandboxServer) Unary(ctx context.Context, request *sandboxpb.UnaryRequest) (*sandboxpb.UnaryResponse, error) {
	return &sandboxpb.UnaryResponse{ Response: "hello" }, nil
}

func (s SandboxServer) ServerStream(request *sandboxpb.UnaryRequest, server sandboxpb.Sandbox_ServerStreamServer) error {
	for i := 0; i < 60; i++ {
		res := &sandboxpb.UnaryResponse{ Response: fmt.Sprintf("%d", i)}
		err := server.Send(res)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s SandboxServer) ClientStream(server sandboxpb.Sandbox_ClientStreamServer) error {
	log.Println("client streaming...")
	for i := 0; i < 60; i++ {
		req, err := server.Recv()
		if err != nil {
			return err
		}
		log.Println(req.Message)
	}
	return nil
}

func (s SandboxServer) BidirectionalStream(server sandboxpb.Sandbox_BidirectionalStreamServer) error {
	log.Println("bidirectional streaming...")
	for i := 0; i < 60; i++ {
		req, err := server.Recv()
		if err != nil {
			return err
		}
		log.Println("client msg:", req.Message)

		res := &sandboxpb.UnaryResponse{
			Response: fmt.Sprintf("%c", rune(65 + i)),
		}
		err = server.Send(res)
	}
	return nil
}
