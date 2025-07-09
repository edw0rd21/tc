package cmd

import (
	"fmt"
	"os"
	"tc/internal/clipboard"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tc",
	Short: "Terminal Clipboard - A CLI clipboard manager for Mac",
	Long:  `Terminal Clipboard (tc) is a CLI tool that keeps track of your clipboard history and allows you to access previous clipboard items.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Default behavior: show last 2 items
		showLastItems(2)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func showLastItems(n int) {
	manager, err := clipboard.NewManager()
	if err != nil {
		fmt.Printf("Error initializing clipboard manager: %v\n", err)
		return
	}
	
	items, err := manager.GetLastItems(n)
	if err != nil {
		fmt.Printf("Error getting clipboard history: %v\n", err)
		return
	}
	
	if len(items) == 0 {
		fmt.Println("No clipboard history found.")
		return
	}
	
	fmt.Printf("Last %d clipboard items:\n\n", len(items))
	for i, item := range items {
		fmt.Println(manager.FormatItem(item, i))
	}
	
	fmt.Println("\nUse 'tc copy <number>' to copy an item back to clipboard")
}