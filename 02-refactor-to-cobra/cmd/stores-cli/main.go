package main

import (
	"github.com/CodelyTV/golang-examples/02-refactor-to-cobra/internal/cli/stores"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "stores-cli"}
	rootCmd.AddCommand(stores.InitStoresCmd())
	rootCmd.Execute()
}
