package function1

import (
	"cloud.google.com/go/pubsub"
	"context"
	_ "github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"github.com/dev2choiz/hello/pkg/app_wire"
	f1 "github.com/dev2choiz/hello/pkg/function1"
	"github.com/dev2choiz/hello/pkg/logger"
	"go.uber.org/zap"
)

type PubSubMessage struct {
	Data []byte `json:"data"`
}

func init() {
	app_wire.InitializeLogger()
}

func Execute(ctx context.Context, m pubsub.Message) error {
	logger.Info("received msg:", zap.String("data", string(m.Data)))
	name := string(m.Data)
	if name == "" {
		name = "World"
	}

	return f1.DoFunction1(name)
}
