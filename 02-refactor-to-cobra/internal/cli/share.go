package cli

import "github.com/spf13/cobra"

// CobraFn function definion of run cobra command
type CobraFn func(cmd *cobra.Command, args []string)

// ID flag definition
const idFlag = "id"
