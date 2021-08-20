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
		<-time.After(time.Duration(1) * time.Second)
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
			log.Println(req.Message)
		}
	}
}

func (s SandboxServer) BidirectionalStream(server sandboxpb.Sandbox_BidirectionalStreamServer) error {
	writeErr := make(chan error)
	readErr := make(chan error)
	end := make(chan int)
	_ = end
	go func() {
		for i := 0; i < 25; i++ {
			res := &sandboxpb.UnaryResponse{
				Response: fmt.Sprintf("%c", rune(65 + i)),
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
				log.Println("received:", req.Message)
			}
		}
	}()

	for i := 0; i < 2; i++ {
		select {
		case err:= <- writeErr:
			if err != nil {
				return err
			}
			log.Println("end server stream")
		case err:= <- readErr:
			if err != nil {
				return err
			}
			log.Println("end client stream")
		}
	}

	return nil
}
