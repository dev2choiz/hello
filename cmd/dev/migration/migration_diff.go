package migration

import (
	"github.com/dev2choiz/hello/pkg/app_wire"
	"github.com/dev2choiz/hello/pkg/pg"
	"github.com/dev2choiz/hello/pkg/pg_migration/diff"
	"github.com/spf13/cobra"
)

var migrationDiffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Generate a diff migration",
	Long:  "Generate a diff migration",
	Run:   runMigrationDiff,
}

func runMigrationDiff(cmd *cobra.Command, args []string) {
	app_wire.InitApp()
	db := pg.GetDB()
	err := diff.Generate(db)
	if err != nil {
		panic(err)
	}
}
