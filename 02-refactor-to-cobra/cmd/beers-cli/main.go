package main

import (
	"github.com/Codely/golang-examples/02-refactor-to-cobra/internal/cli"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(cli.InitBeersCmd())
	rootCmd.AddCommand(cli.InitStoresCommand())
	rootCmd.Execute()
}
