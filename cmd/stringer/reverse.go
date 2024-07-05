package stringer

import (
	"fmt"

	"github.com/alanjose10/stringer-cli/pkg/stringer"
	"github.com/spf13/cobra"
)

var reverseCmd = &cobra.Command{
	Use:     "reverse",
	Aliases: []string{"rev", "r"},
	Short:   "Reverse a string",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res := stringer.Reverse(args[0])
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(reverseCmd)
}
