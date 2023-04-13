package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

// CobraFn function definion of run cobra command
type CobraAFn func(cmd *cobra.Command, args []string)

var stores = map[string]string{
	"01DC9ZAPGKEQJS4P4A48EG3P43": "Mercadona",
	"01DC9ZB23EW0J0ARAER09SJDKC": "Carrefour",
	"01DC9ZB89V1PQD977ZE6QXSQHH": "Alcampo",
}

const idAFlag = "id"

// InitStoresCmd initialize stores command
func InitStoresCmd() *cobra.Command {
	storesCmd := &cobra.Command{
		Use:   "stores",
		Short: "Print data about beer stores",
		Run:   runStoresFn(),
	}

	storesCmd.Flags().StringP(idFlag, "i", "", "id of the beer store")

	return storesCmd
}

func runStoresFn() CobraAFn {
	return func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString(idAFlag)

		if id != "" {
			fmt.Println(stores[id])
		} else {
			fmt.Println(stores)
		}
	}
}
