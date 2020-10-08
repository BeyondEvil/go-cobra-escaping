package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	withCmd = &cobra.Command{
		Use:   "with <message>",
		Short: `Print argument with escaping`,
		Long:  `Print an argument passed through cobra with escaping`,
		Run: func(cmd *cobra.Command, args []string) {
			with(args[0])
		},
	}
)

func init() {
	rootCmd.AddCommand(withCmd)
}

func with(argument string) {
	fmt.Printf("Raw string: %q\n", argument)
	fmt.Printf("Result with cobra escaping: %s\n", argument)
}
