package function1

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	f1 "github.com/dev2choiz/hello/pkg/function1"
)

type PubSubMessage struct {
	Data []byte `json:"data"`
}

func Function1(ctx context.Context, m pubsub.Message) error {
	fmt.Println("received msg:")
	fmt.Println("Data", m.Data)
	name := string(m.Data)
	if name == "" {
		name = "World"
	}

	return f1.DoFunction1(name)
}
