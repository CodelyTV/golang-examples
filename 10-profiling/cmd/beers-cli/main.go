package main

import (
	"flag"

	beerscli "github.com/CodelyTV/golang-examples/10-profiling/internal"
	"github.com/CodelyTV/golang-examples/10-profiling/internal/cli"
	"github.com/CodelyTV/golang-examples/10-profiling/internal/fetching"
	"github.com/CodelyTV/golang-examples/10-profiling/internal/storage/csv"
	"github.com/CodelyTV/golang-examples/10-profiling/internal/storage/ontario"
	"github.com/spf13/cobra"
)

func main() {
	// CPU profiling code starts here
	//f, _ := os.Create("beers.cpu.prof")
	//defer f.Close()
	//pprof.StartCPUProfile(f)
	//defer pprof.StopCPUProfile()
	// CPU profiling code ends here

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

	// Memory profiling code starts here
	//f2, _ := os.Create("beers.mem.prof")
	//defer f2.Close()
	//pprof.WriteHeapProfile(f2)
	// Memory profiling code ends here
}
