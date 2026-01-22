package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCommand = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"d"},
	Short:   "Delete an item in your list",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is where I would delete some shit")

	},
}

func init() {
	rootCmd.AddCommand(deleteCommand)
}
