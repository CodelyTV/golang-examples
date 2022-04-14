package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// CobraFn function definion of run cobra command
type CobraFn func(cmd *cobra.Command, args []string)

var beers = map[string]string{
	"01D9X58E7NPXX5MVCR9QN794CH": "Mad Jack Mixer",
	"01D9X5BQ5X48XMMVZ2F2G3R5MS": "Keystone Ice",
	"01D9X5CVS1M9VR5ZD627XDF6ND": "Belgian Moon",
}

var stores = map[string]string{
	"01DC9ZAPGKEQJS4P4A48EG3P43": "Mercadona",
	"01DC9ZB23EW0J0ARAER09SJDKC": "Carrefour",
	"01DC9ZB89V1PQD977ZE6QXSQHH": "Alcampo",
}

const idFlag = "id"
const storeIdFlag = "store"

// InitBeersCmd initialize beers command
func InitBeersCmd() *cobra.Command {
	beersCmd := &cobra.Command{
		Use:   "beers",
		Short: "Print data about beers",
		Run:   runBeersFn(),
	}

	beersCmd.Flags().StringP(idFlag, "i", "", "id of the beer")

	return beersCmd
}

// InitStoresCmd initialize stores command
func InitStoresCmd() *cobra.Command {
	storeCmd := &cobra.Command{
		Use:   "stores",
		Short: "Print data about stores that sell beers",
		Run:   runStoresFn(),
	}

	storeCmd.Flags().StringP(storeIdFlag, "s", "", "id of the store")

	return storeCmd
}

func runBeersFn() CobraFn {
	return func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString(idFlag)

		if id != "" {
			fmt.Println(beers[id])
		} else {
			fmt.Println(beers)
		}
	}
}

func runStoresFn() CobraFn {
	return func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString(storeIdFlag)

		if id != "" {
			fmt.Println(stores[id])
		} else {
			fmt.Println(stores)
		}
	}
}
