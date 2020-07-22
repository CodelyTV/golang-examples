package markets

import (
	"fmt"

	"github.com/spf13/cobra"
)

// CobraFn function definion of run cobra command
type CobraFn func(cmd *cobra.Command, args []string)

var markets = map[string]string{
	"01DC9ZAPGKEQJS4P4A48EG3P43": "Mercadona",
	"01DC9ZB23EW0J0ARAER09SJDKC": "Carrefour",
	"01DC9ZB89V1PQD977ZE6QXSQHH": "Alcampo",
}

const idFlag = "id"

// InitMarketsCmd initialize beers command
func InitMarketsCmd() *cobra.Command {
	marketsCmd := &cobra.Command{
		Use:   "markets",
		Short: "Print data about markets",
		Run:   runMarketsFn(),
	}

	marketsCmd.Flags().StringP(idFlag, "i", "", "id of the market")

	return marketsCmd
}

func runMarketsFn() CobraFn {
	return func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString(idFlag)

		if id != "" {
			fmt.Println(markets[id])
		} else {
			fmt.Println(markets)
		}
	}
}
