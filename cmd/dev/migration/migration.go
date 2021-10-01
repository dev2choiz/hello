package migration

import (
	"github.com/dev2choiz/hello/pkg/config"
	"github.com/dev2choiz/hello/pkg/pg_migration"
	"github.com/spf13/cobra"
)

var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Execute a pg migration down",
	Long: "Execute a pg migration down",
	Run: func(cmd *cobra.Command, args []string) {
		executeMigration([]string{"down"})
	},
}

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Execute a pg migration up",
	Long: "Execute a pg migration up",
	Run: func(cmd *cobra.Command, args []string) {
		executeMigration([]string{"up"})
	},
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Execute a pg migration init",
	Long: "Execute a pg migration init",
	Run: func(cmd *cobra.Command, args []string) {
		executeMigration([]string{"init"})
	},
}

func executeMigration(params []string) {
	conf := config.Conf
	//dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient debug *.go migration up
	err := pg_migration.Migrate(params, conf)
	if err != nil {
		panic(err)
	}
}