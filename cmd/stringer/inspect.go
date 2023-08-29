package stringer

import (
	"cliapp/pkg/stringer"
	"fmt"
	"github.com/spf13/cobra"
)

var countCmd = &cobra.Command{
	Use:     "count",
	Aliases: []string{"cou"},
	Short:   "gives the number of elements in a string",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res := stringer.Count(args[0])
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(countCmd)
}
