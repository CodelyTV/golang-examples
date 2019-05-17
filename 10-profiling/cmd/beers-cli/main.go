package main

import (
	"flag"
	"os"
	"runtime/pprof"

	beerscli "github.com/CodelyTV/golang-introduction/10-profiling/internal"
	"github.com/CodelyTV/golang-introduction/10-profiling/internal/cli"
	"github.com/CodelyTV/golang-introduction/10-profiling/internal/fetching"
	"github.com/CodelyTV/golang-introduction/10-profiling/internal/storage/csv"
	"github.com/CodelyTV/golang-introduction/10-profiling/internal/storage/ontario"
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

	fetchingService := fetching.NewService(repo)

	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(cli.InitBeersCmd(fetchingService))
	rootCmd.Execute()

	f, _ := os.Create("beers.mem.prof")
	defer f.Close()
	pprof.WriteHeapProfile(f)
}
