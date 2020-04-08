package clistores

import (
	"fmt"
	"log"
	"strconv"

	"github.com/CodelyTV/golang-examples/08-automated_tests/internal/fetching"

	"github.com/spf13/cobra"
)

// CobraFn function definion of run cobra command
type CobraFn func(cmd *cobra.Command, args []string)

const idFlag = "id"

// InitStoresCmd initialize stores command
func InitStoresCmd(service fetching.ServiceStores) *cobra.Command {
	storesCmd := &cobra.Command{
		Use:   "stores",
		Short: "Print data about stores",
		Run:   runStoresFn(service),
	}

	storesCmd.Flags().StringP(idFlag, "i", "", "id of the store")

	return storesCmd
}

func runStoresFn(service fetching.ServiceStores) CobraFn {
	return func(cmd *cobra.Command, args []string) {

		id, _ := cmd.Flags().GetString(idFlag)
		if id != "" {
			i, _ := strconv.Atoi(id)
			store, err := service.FetchByID(i)
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
