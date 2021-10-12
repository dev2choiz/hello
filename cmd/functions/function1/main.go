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

func init() {
	app_wire.InitApp()
}

// Execute the entrypoint executed in the cloud function
func Execute(ctx context.Context, m pubsub.Message) error {
	logger.Info("received msg:", zap.String("data", string(m.Data)))
	name := string(m.Data)
	if name == "" {
		name = "World"
	}

	if err := f1.DoFunction1(name); err != nil {
		logger.Error(err.Error())
	}
	return nil // no retry
}
