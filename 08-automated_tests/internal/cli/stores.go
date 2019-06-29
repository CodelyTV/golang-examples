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

// InitStoresCmd initialize stores command
func InitStoresCmd(service fetching.StoresService) *cobra.Command {
	storesCmd := &cobra.Command{
		Use:   "stores",
		Short: "Print data about stores",
		Run:   runStoresFn(service),
	}

	storesCmd.Flags().StringP(idStoreFlag, "i", "", "id of the store")

	return storesCmd
}

func runStoresFn(service fetching.StoresService) CobraFn {
	return func(cmd *cobra.Command, args []string) {

		id, _ := cmd.Flags().GetString(idStoreFlag)
		if id != "" {
			i, _ := strconv.Atoi(id)
			store, err := service.FetchStoreByID(i)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(store)
			return
		}

		stores, err := service.FetchStores()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(stores)
	}
}
