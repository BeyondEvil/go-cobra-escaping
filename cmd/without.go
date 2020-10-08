package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var (
	withoutCmd = &cobra.Command{
		Use:   "without <message>",
		Short: `Print argument without escaping`,
		Long:  `Print an argument passed through cobra without escaping`,
		Run: func(cmd *cobra.Command, args []string) {
			without(args[0])
		},
	}
)

func init() {
	rootCmd.AddCommand(withoutCmd)
}

func without(argument string) {
	fmt.Printf("Raw string: %q\n", argument)
	argument = strings.ReplaceAll(argument, "\\n", "\n")
	fmt.Printf("Result without cobra escaping: %s\n", argument)
}
