package commands

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "Todo command can manage your todos.",
	Long: `Todo command is intended to help you to
manage your todos in terminal.

Todo can create, delete, fetch and update your todos.
I created this simple CLI tool in order to manage my todos.
I managed to have very short and descriptive todos,
so I wasn't happy with the idea of having very "big" application for 3-4 one line notes.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(newCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(finishCmd)
	rootCmd.AddCommand(updateCmd)
}
