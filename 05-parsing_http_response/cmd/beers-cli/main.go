package main

import (
	"flag"

	"github.com/CodelyTV/golang-examples/05-parsing_http_response/internal/storage/ontario"

	beerscli "github.com/CodelyTV/golang-examples/05-parsing_http_response/internal"
	"github.com/CodelyTV/golang-examples/05-parsing_http_response/internal/cli"
	"github.com/CodelyTV/golang-examples/05-parsing_http_response/internal/storage/csv"
	"github.com/spf13/cobra"
)

func main() {

	csvData := flag.Bool("csv", false, "load data from csv")
	flag.Parse()

	var repo beerscli.BeerRepo
	repo = csv.NewRepository()

	if !*csvData {
		repo = ontario.NewOntarioRepository()
	}

	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(cli.InitBeersCmd(repo))
	rootCmd.Execute()
}
