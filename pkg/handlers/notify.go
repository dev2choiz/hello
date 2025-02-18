package handlers

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"github.com/dev2choiz/hello/pkg/config"
	"github.com/dev2choiz/hello/pkg/protobuf/notifypb"
)

type NotifyServer struct {
	notifypb.UnimplementedNotifyServer
}

// Status handler to call a cloud function using pubsub
func (n NotifyServer) Status(ctx context.Context, req *notifypb.Function1Request) (*notifypb.Function1Response, error) {
	res := &notifypb.Function1Response{}

	client, err := pubsub.NewClient(ctx, config.Conf.GcpProjectName)
	if err != nil {
		return nil, fmt.Errorf("can't create pubsub client:" + err.Error())
	}
	defer client.Close()

	t := client.Topic("hello-function1")
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(req.Name),
	})

	<-result.Ready()
	res.PsId, err = result.Get(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}
