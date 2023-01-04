package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

const idFlag = "id"
const shortIdFlag = "i"

type CobraFn func(cmd *cobra.Command, args []string)

func runCommandFn(values map[string]string) CobraFn {
	return func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString(idFlag)

		if id != "" {
			fmt.Println(values[id])
		} else {
			fmt.Println(values)
		}
	}
}
