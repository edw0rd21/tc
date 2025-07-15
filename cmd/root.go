package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tc",
	Short: "Terminal Clipboard - A CLI clipboard manager for Mac",
	Long: `Terminal Clipboard (tc) is a CLI tool that monitors your clipboard history and lets you interact with past clipboard content.

Available commands:
  list         Show recent clipboard items
  copy <n>     Copy a specific item from history back to the clipboard
  clear        Clear clipboard history
  daemon       Start background clipboard watcher

Examples:
  tc list            Show last 10 items
  tc list 3          Show 3rd item from history
  tc copy 2          Copy item #2 to clipboard
  tc clear           Erase clipboard history
  tc daemon          Start background watcher`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
