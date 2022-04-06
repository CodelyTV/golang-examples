package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// CobraStoreFn function definion of run cobra command
type CobraStoreFn func(cmd *cobra.Command, args []string)

const idStoreFlag = "id"

var stores = map[string]string{
	"01DC9ZAPGKEQJS4P4A48EG3P43": "Mercadona",
	"01DC9ZB23EW0J0ARAER09SJDKC": "Carrefour",
	"01DC9ZB89V1PQD977ZE6QXSQHH": "Alcampo",
}

func InitStoresCmd() *cobra.Command {
	storesCmd := &cobra.Command{
		Use:   "stores",
		Short: "Print data about stores",
		Run:   runStoresFn(),
	}
	storesCmd.Flags().StringP(idStoreFlag, "i", "", "id of the store")
	return storesCmd
}

func runStoresFn() CobraStoreFn {
	return func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString(idStoreFlag)

		if id != "" {
			fmt.Println(stores[id])
		} else {
			fmt.Println(stores)
		}
	}
}
