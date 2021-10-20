package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

type CobraStoreFn func(cmd *cobra.Command, args []string)

var stores = map[string]string{
	"01DC9ZAPGKEQJS4P4A48EG3P43": "Mercadona",
	"01DC9ZB23EW0J0ARAER09SJDKC": "Carrefour",
	"01DC9ZB89V1PQD977ZE6QXSQHH": "Alcampo",
}

const idStoreFlag = "store"

func InitStoresCmd() *cobra.Command {
	storesCmd := &cobra.Command{
		Use:   "stores",
		Short: "Print data about stores",
		Run:   runStoresFn(),
	}

	storesCmd.Flags().StringP(idStoreFlag, "s", "", "id of the store")

	return storesCmd
}

func runStoresFn() CobraStoreFn {
	return func(cmd *cobra.Command, args []string) {
		store, _ := cmd.Flags().GetString(idStoreFlag)

		if store != "" {
			fmt.Println(stores[store])
		} else {
			fmt.Println(stores)
		}
	}
}
