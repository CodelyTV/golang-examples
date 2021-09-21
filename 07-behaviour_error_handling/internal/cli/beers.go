package cli

import (
	"fmt"
	"log"
	"strconv"

	"github.com/CodelyTV/golang-examples/07-behaviour_error_handling/internal/errors"

	beerscli "github.com/CodelyTV/golang-examples/07-behaviour_error_handling/internal"
	"github.com/spf13/cobra"
)

// CobraFn function definion of run cobra command
type CobraFn func(cmd *cobra.Command, args []string)

const idFlag = "id"

// InitBeersCmd initialize beers command
func InitBeersCmd(repoAPI beerscli.BeerRepo, repoCSV beerscli.BeerRepo) *cobra.Command {
	beersCmd := &cobra.Command{
		Use:   "beers",
		Short: "Print data about beers",
		Run:   runBeersFn(repoAPI, repoCSV),
	}

	beersCmd.Flags().StringP(idFlag, "i", "", "id of the beer")
	beersCmd.Flags().Bool("csv", false, "load data from csv")

	return beersCmd
}

func runBeersFn(repoAPI beerscli.BeerRepo, repoCSV beerscli.BeerRepo) CobraFn {
	return func(cmd *cobra.Command, args []string) {
		var repository beerscli.BeerRepo
		useCSV, _ := cmd.Flags().GetBool("csv")
		if useCSV {
			repository = repoAPI
		} else {
			repository = repoCSV
		}
		beers, err := repository.GetBeers()
		if errors.IsDataUnreacheable(err) {
			log.Printf("https://github.com/caleeli/golang-examples.git/issues/new?assignees=&labels=&template=dataunreachable.md&title=")
		}
		if errors.IsLoadCSVError(err) {
			log.Printf("https://github.com/caleeli/golang-examples.git/issues/new?assignees=&labels=&template=loadcsv.md&title=")
		}
		if err != nil {
			log.Fatal(err)
		}

		id, _ := cmd.Flags().GetString(idFlag)

		if id != "" {
			i, _ := strconv.Atoi(id)
			for _, beer := range beers {
				if beer.ProductID == i {
					fmt.Println(beer)
					return
				}
			}
		} else {
			fmt.Println(beers)
		}
	}
}
