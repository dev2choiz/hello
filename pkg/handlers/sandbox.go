package handlers

import (
	"context"
	"fmt"
	"github.com/dev2choiz/hello/pkg/protobuf/sandboxpb"
	"io"
	"log"
	"time"
)

type SandboxServer struct {
	sandboxpb.UnimplementedSandboxServer
}

func (s SandboxServer) Unary(ctx context.Context, req *sandboxpb.UnaryRequest) (*sandboxpb.UnaryResponse, error) {
	name := req.GetName()
	if name == "" {
		name = "sir"
	}
	return &sandboxpb.UnaryResponse{Response: "hello " + name}, nil
}

func (s SandboxServer) ServerStream(req *sandboxpb.ServerStreamRequest, server sandboxpb.Sandbox_ServerStreamServer) error {
	ms := int(req.MsPerResponse)
	nb := int(req.Number)
	if nb == 0 {
		nb = 10
	}
	for i := 0; i < nb; i++ {
		res := &sandboxpb.ServerStreamResponse{Message: fmt.Sprintf("%d", i)}
		err := server.Send(res)
		if err != nil {
			return err
		}
		if ms != 0 && i < nb-1 {
			<-time.After(time.Duration(ms) * time.Millisecond)
		}
	}
	return nil
}

func (s SandboxServer) ClientStream(server sandboxpb.Sandbox_ClientStreamServer) error {
	for {
		select {
		case <-server.Context().Done():
			return nil
		default:
			req, err := server.Recv()
			if err == io.EOF {
				log.Println("end client stream")
				return nil
			}
			if err != nil {
				return err
			}
			log.Println(req.Name)
		}
	}
}

func (s SandboxServer) BidirectionalStream(server sandboxpb.Sandbox_BidirectionalStreamServer) error {
	writeErr := make(chan error)
	readErr := make(chan error)
	go func() {
		for i := 0; i < 25; i++ {
			res := &sandboxpb.UnaryResponse{
				Response: fmt.Sprintf("%c", rune(65+i)),
			}
			err := server.Send(res)
			if err != nil {
				writeErr <- err
				return
			}
			<-time.After(time.Duration(1) * time.Second)
		}
		writeErr <- nil
		return
	}()

	go func() {
		for {
			select {
			case <-server.Context().Done():
				readErr <- server.Context().Err()
				return
			default:
				req, err := server.Recv()
				if err == io.EOF {
					readErr <- nil
					return
				}
				if err != nil {
					readErr <- err
					return
				}
				log.Println("received:", req.Name)
			}
		}
	}()

	for i := 0; i < 2; i++ {
		select {
		case err := <-writeErr:
			if err != nil {
				return err
			}
			log.Println("end server stream")
		case err := <-readErr:
			if err != nil {
				return err
			}
			log.Println("end client stream")
		}
	}

	return nil
}
