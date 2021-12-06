package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// CobraFn function definion of run cobra command
type CobraFn func(cmd *cobra.Command, args []string)

var shops = map[string]string{
	"01DC9ZAPGKEQJS4P4A48EG3P43": "Mercadona",
	"01DC9ZB23EW0J0ARAER09SJDKC": "Carrefour",
	"01DC9ZB89V1PQD977ZE6QXSQHH": "Alcampo",
}

const idFlag = "id"

// InitshopsCmd initialize beers command
func InitShopsCmd() *cobra.Command {
	shopsCmd := &cobra.Command{
		Use:   "shops",
		Short: "Print data about shops",
		Run:   runShopsFn(),
	}

	shopsCmd.Flags().StringP(idFlag, "i", "", "id of the shop")

	return shopsCmd
}

func runShopsFn() CobraFn {
	return func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString(idFlag)

		if id != "" {
			fmt.Println(shops[id])
		} else {
			fmt.Println(shops)
		}
	}
}
