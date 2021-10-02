package migration

import (
	"cloud.google.com/go/pubsub"
	"context"
	_ "github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"github.com/dev2choiz/hello/pkg/app_wire"
	"github.com/dev2choiz/hello/pkg/config"
	"github.com/dev2choiz/hello/pkg/pg_migration"
)

func init() {
	app_wire.InitializeLogger()
	app_wire.InitializePostgres()
}

// Execute the entrypoint executed in the cloud function
func Execute(ctx context.Context, m pubsub.Message) error {
	conf := config.Conf
	cmd := string(m.Data)
	// cmd should be "init" | "up" | "down"
	params := []string{cmd}
	return pg_migration.Migrate(params, conf)
}
