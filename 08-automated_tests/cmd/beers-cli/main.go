package main

import (
	"flag"

	"github.com/CodelyTV/golang-introduction/08-automated_tests/internal/fetching"

	"github.com/CodelyTV/golang-introduction/08-automated_tests/internal/storage/ontario"

	beerscli "github.com/CodelyTV/golang-introduction/08-automated_tests/internal"
	"github.com/CodelyTV/golang-introduction/08-automated_tests/internal/cli"
	"github.com/CodelyTV/golang-introduction/08-automated_tests/internal/storage/csv"
	"github.com/spf13/cobra"
)

func main() {

	csvData := flag.Bool("csv", false, "load data from csv")
	flag.Parse()
	var repo beerscli.BeerRepo
	var storeRepo beerscli.StoreRepo
	repo = csv.NewRepository()

	if !*csvData {
		repo, storeRepo = ontario.NewOntarioRepository()
	}

	fetchingService := fetching.NewService(repo)
	fetchingStoresService := fetching.NewStoreService(storeRepo)

	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(cli.InitBeersCmd(fetchingService))
	rootCmd.AddCommand(cli.InitStoresCmd(fetchingStoresService))
	rootCmd.Execute()
}
