package cmd

import (
	"fmt"
	"strconv"

	"github.com/edw0rd21/tc/internal/clipboard"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list [count]",
	Short: "List clipboard history items",
	Long:  `Displays the most recent clipboard history items. Optionally specify how many items to show.`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		count, _ := cmd.Flags().GetInt("count")

		// If a positional argument is provided, override the flag
		if len(args) > 0 {
			userCount, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Printf("tc list: Invalid count: %s\n", args[0])
				return
			}
			if userCount < 1 {
				fmt.Println("tc list: Count must be 1 or greater")
				return
			}
			count = userCount // Override flag with positional
		}

		manager, err := clipboard.NewManager()
		if err != nil {
			fmt.Printf("tc list: Error initializing clipboard manager: %v\n", err)
			return
		}

		items, err := manager.GetLastItems(count)
		if err != nil {
			fmt.Printf("tc list: Error retrieving clipboard history: %v\n", err)
			return
		}

		if len(items) == 0 {
			fmt.Println("No clipboard history found.")
			return
		}

		for i, item := range items {
			fmt.Println(manager.FormatItem(item, i))
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().IntP("count", "n", 10, "Number of clipboard items to list (alternative to positional argument)")
}
