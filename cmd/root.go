package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func checkError(e error) {
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
		os.Exit(1)
	}
}

func PrintUsageMsg(command string, msgCase string) {
	switch msgCase {
	case "add_none":
		fmt.Println("Missing task description")
	case "add_to_many":
		fmt.Println("Too many arguments passed to the \"add\" command")

	}

	switch command {
	case "add":
		fmt.Println("Usage: killahtask add \"my description\"")
	}
}

var rootCmd = &cobra.Command{
	Use:   "killahtask",
	Short: "Killah Task is a todo CLI tool.",
	Long:  `A todo task tool that does things`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("No commands were passed to the killah...see below")
			cmd.Help()
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	checkError(err)
}
