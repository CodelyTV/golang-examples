package cli

import (
	"github.com/spf13/cobra"
)

var stores = map[string]string{
	"01DC9ZAPGKEQJS4P4A48EG3P43": "Mercadona",
	"01DC9ZB23EW0J0ARAER09SJDKC": "Carrefour",
	"01DC9ZB89V1PQD977ZE6QXSQHH": "Alcampo",
}

// InitStoresCmd initialize stores command
func InitStoresCmd() *cobra.Command {
	beersCmd := &cobra.Command{
		Use:   "stores",
		Short: "Print data about stores",
		Run:   listItems(stores),
	}

	beersCmd.Flags().StringP(idFlag, "i", "", "id of the store")

	return beersCmd
}
