package main

import (
	"github.com/spf13/cobra"

	"github.com/CodelyTV/golang-examples/02-refactor-to-cobra/internal/cli"
)

func main() {
	rootCmd := &cobra.Command{Use: "beers-cli"}

	rootCmd.AddCommand(cli.NewCommand("beers"))

	rootCmd.AddCommand(cli.NewCommand("stores"))

	rootCmd.Execute()
}
