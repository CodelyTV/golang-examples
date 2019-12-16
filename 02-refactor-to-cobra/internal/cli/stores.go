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
const storeFlag = "store"

// InitStoresCmd initialize stores command
func InitStoresCmd() *cobra.Command {
	storesCmd := &cobra.Command{
		Use:   "stores",
		Short: "Print where i can buy my beer",
		Run:   runStoressFn(),
	}

	storesCmd.Flags().StringP(storeFlag, "s", "", "store of the beer")

	return storesCmd
}

func runStoressFn() CobraFn {
	return func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString(storeFlag)
		
		if id != "" {
			fmt.Println(stores[id])
		} else {
			fmt.Println(stores)
		}
	}
}
