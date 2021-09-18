package app

import (
	"github.com/dev2choiz/hello/pkg/app_wire"
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

func init() {
	app_wire.InitApp()
}

func Execute() error {
	rootCmd.AddCommand(helloApiCmd)
	rootCmd.AddCommand(helloSvcCmd)
	return rootCmd.Execute()
}