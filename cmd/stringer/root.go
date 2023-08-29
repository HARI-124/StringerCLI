package stringer

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "stringer",
	Short: "stringer - a simple cli to transform and inspect strings",
	Long:  `stringer is a super fancy cli`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "whoops there was an error  while executing your cli %s", err)
		os.Exit(1)
	}
}
