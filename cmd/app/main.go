package app

import (
	"github.com/dev2choiz/hello/pkg/app_wire"
	"github.com/spf13/cobra"
)

var usage = `hello CLI`

var rootCmd = &cobra.Command{
	Short: "",
	Long: usage,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {
	app_wire.InitializeLogger()
}

func Execute() error {
	rootCmd.AddCommand(helloApiCmd)
	rootCmd.AddCommand(helloSvcCmd)
	return rootCmd.Execute()
}