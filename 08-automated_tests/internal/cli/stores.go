package cli

import (
	"fmt"
	"log"
	"strconv"

	"github.com/CodelyTV/golang-introduction/08-automated_tests/internal/fetching"

	"github.com/spf13/cobra"
)

// CobraFn function definion of run cobra command
type CobraFn func(cmd *cobra.Command, args []string)

const idStoreFlag = "id"

// InitBeersCmd initialize beers command
func InitStoresCmd(service fetching.Service) *cobra.Command {
	beersCmd := &cobra.Command{
		Use:   "stores",
		Short: "Print data about stores",
		Run:   runStoresFn(service),
	}

	beersCmd.Flags().StringP(idStoreFlag, "i", "", "id of the store")

	return beersCmd
}

func runStoresFn(service fetching.Service) CobraFn {
	return func(cmd *cobra.Command, args []string) {

		id, _ := cmd.Flags().GetString(idStoreFlag)
		if id != "" {
			i, _ := strconv.Atoi(id)
			beer, err := service.FetchByID(i)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(beer)
			return
		}

		beers, err := service.FetchBeers()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(beers)
	}
}
