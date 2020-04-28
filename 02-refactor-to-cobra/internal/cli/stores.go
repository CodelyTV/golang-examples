package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var stores = map[string]string{
	"01DC9ZAPGKEQJS4P4A48EG3P43": "Mercadona",
	"01DC9ZB23EW0J0ARAER09SJDKC": "Carrefour",
	"01DC9ZB89V1PQD977ZE6QXSQHH": "Alcampo",
}

func InitStoresCommand() *cobra.Command {
	command := &cobra.Command{
		Use: "stores",
		Short: "Print store about beers",
		Run: runStoresFn(),
	}

	command.Flags().StringP("id", "i","", "id of a store" )
	return command
}

func runStoresFn() func(cmd *cobra.Command, args []string) {
	var printStoreByBeerId = func(cmd *cobra.Command, args []string) {
		var beerId, _ = cmd.Flags().GetString("id")

		if beerId != "" {
			fmt.Println(stores[beerId])
		} else {
			fmt.Println(stores)
		}
	}

	return printStoreByBeerId
}