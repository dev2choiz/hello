package app

import (
	"github.com/spf13/cobra"
)

var rootDesc = `hello CLI`

var rootCmd = &cobra.Command{
	Short: rootDesc,
	Long: rootDesc,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

// Execute initialize cobra sub commands
func Execute() error {
	rootCmd.AddCommand(helloApiCmd)
	rootCmd.AddCommand(helloSvcCmd)
	return rootCmd.Execute()
}
