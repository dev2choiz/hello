package function1

import (
	"cloud.google.com/go/pubsub"
	"context"
	"github.com/dev2choiz/hello/internal/logger"
	f1 "github.com/dev2choiz/hello/pkg/function1"
	"go.uber.org/zap"
)

type PubSubMessage struct {
	Data []byte `json:"data"`
}

func Function1(ctx context.Context, m pubsub.Message) error {
	logger.Info("received msg:", zap.String("data", string(m.Data)))
	name := string(m.Data)
	if name == "" {
		name = "World"
	}

	return f1.DoFunction1(name)
}
