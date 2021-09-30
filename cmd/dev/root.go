package main

import (
	"github.com/dev2choiz/hello/cmd/dev/migration"
	"github.com/dev2choiz/hello/pkg/app_wire"
	"github.com/spf13/cobra"
)

var rootDesc = `hello dev CLI`

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

func main() {
	rootCmd.AddCommand(migration.RootCmd)
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}