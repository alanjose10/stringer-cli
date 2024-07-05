package stringer

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "stringer",
	Short: "stringer - a simple CLI to transform and inspect strings",
	Long: `stringer is a super fancy CLI (kidding)
	
One can use stringer to modify or inspect strings straight from the terminal`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Woops. There was an error while executing the CLI '%s'", err)
		os.Exit(1)
	}
}
