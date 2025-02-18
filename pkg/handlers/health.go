package handlers

import (
	"context"
	"github.com/dev2choiz/hello/pkg/logger"
	"github.com/dev2choiz/hello/pkg/models"
	"github.com/dev2choiz/hello/pkg/pg"
	"github.com/dev2choiz/hello/pkg/protobuf/healthpb"
	"github.com/dev2choiz/hello/pkg/protobuf/pingpb"
	"github.com/dev2choiz/hello/pkg/version"
	"google.golang.org/grpc"
	"os"
)

type HealthServer struct {
	healthpb.UnimplementedHealthServer
}

// Healthz is the handler of Healthz service
func (h HealthServer) Healthz(ctx context.Context, req *healthpb.HealthzRequest) (*healthpb.HealthzResponse, error) {
	return &healthpb.HealthzResponse{Status: "ok"}, nil
}

// Status is the handler of Status service
func (h HealthServer) Status(ctx context.Context, req *healthpb.StatusRequest) (*healthpb.StatusResponse, error) {
	return &healthpb.StatusResponse{Status: "ok"}, nil
}

// Check is a handler checking healthiness of multiple micro-services
func (h HealthServer) Check(ctx context.Context, req *healthpb.CheckServicesRequest) (*healthpb.CheckServicesResponse, error) {
	d := models.Data{ Field1: "default", Field2: "health.check"}
	if db := pg.GetDB().Create(&d); db.Error != nil {
		return nil, db.Error
	}

	res := &healthpb.CheckServicesResponse{}
	data := map[string]string{}
	addSvcData(ctx, data, "svc1", os.Getenv("SVC1_GRPC_BASE_URL"))
	addSvcData(ctx, data, "svc2", os.Getenv("SVC2_GRPC_BASE_URL"))

	res.Status = "ok"
	res.Version = version.Get() // Deprecated: Use wire injection dependencies ?
	res.Data = data

	return res, nil
}

func addSvcData(ctx context.Context, data map[string]string, name, url string) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		logger.Error(err.Error())
		return
	}
	defer conn.Close()
	client := pingpb.NewPingClient(conn)

	req := &pingpb.PingRequest{}
	res, err := client.Ping(ctx, req)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	logger.Infof("%s response: %s", name, res.Response)

	data[name] = res.Response

	return
}
