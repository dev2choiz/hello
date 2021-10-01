package migration

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "migration",
	Short: "Manage pg migrations",
	Long: "Manage pg migrations",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
	RootCmd.AddCommand(upCmd)
	RootCmd.AddCommand(downCmd)
	RootCmd.AddCommand(migrationDiffCmd)
}
