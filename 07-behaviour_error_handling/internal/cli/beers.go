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
func InitBeersCmd(repository beerscli.BeerRepo) *cobra.Command {
	beersCmd := &cobra.Command{
		Use:   "beers",
		Short: "Print data about beers",
		Run:   runBeersFn(repository),
	}

	beersCmd.Flags().StringP(idFlag, "i", "", "id of the beer")

	return beersCmd
}

func runBeersFn(repository beerscli.BeerRepo) CobraFn {
	return func(cmd *cobra.Command, args []string) {
		beers, err := repository.GetBeers()
		if errors.IsFileErrorType(err) {
			log.Fatal(err)
		}
		if errors.IsFormatDataError(err) {
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
