package main

import (
	"github.com/CodelyTV/golang-examples/07-behaviour_error_handling/internal/storage/ontario"

	"github.com/CodelyTV/golang-examples/07-behaviour_error_handling/internal/cli"
	"github.com/CodelyTV/golang-examples/07-behaviour_error_handling/internal/storage/csv"
	"github.com/spf13/cobra"
)

func main() {

	repoAPI := csv.NewRepository()
	repoCSV := ontario.NewOntarioRepository()

	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(cli.InitBeersCmd(repoAPI, repoCSV))
	rootCmd.Execute()
}
