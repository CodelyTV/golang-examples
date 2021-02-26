package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var shops = map[string]string{
	"01D9X58E7NPXX5MVCR9QN794CH": "mc",
	"01D9X5BQ5X48XMMVZ2F2G3R5MS": "ko",
	"01D9X5CVS1M9VR5ZD627XDF6ND": "jju",
}

const idShop = "id"

// InitBeersCmd initialize shops command
func InitShopsCmd() *cobra.Command {
	shopCmd := &cobra.Command{
		Use:   "shops",
		Short: "Print data about shops",
		Run:   runShopsFn(),
	}

	shopCmd.Flags().StringP(idShop, "i", "", "id of the shio")

	return shopCmd
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
