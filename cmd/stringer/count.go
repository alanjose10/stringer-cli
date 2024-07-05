package stringer

import (
	"fmt"

	"github.com/alanjose10/stringer-cli/pkg/stringer"
	"github.com/spf13/cobra"
)

var countCmd = &cobra.Command{
	Use:   "count",
	Short: "count characters in a string",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		i := args[0]
		res := stringer.CountLen(i)

		pl := "s"
		if res == 1 {
			pl = ""
		}
		fmt.Printf("'%s' has %d char%s.\n", i, res, pl)
	},
}

func init() {
	rootCmd.AddCommand(countCmd)
}
