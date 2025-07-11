package cmd

import (
	"fmt"
	"strconv"

	"github.com/edw0rd21/tc/internal/clipboard"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list [index]",
	Short: "List clipboard history items",
	Long: `Displays clipboard history items.

- If a number is passed as an argument (e.g., 'tc list 3'), shows that specific item.
- Use the --count (-n) flag to list the most recent N items (default is 10).`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Check if --count flag was explicitly set
		countFlagUsed := cmd.Flags().Changed("count")
		count, _ := cmd.Flags().GetInt("count")

		manager, err := clipboard.NewManager()
		if err != nil {
			fmt.Printf("tc list: Error initializing clipboard manager: %v\n", err)
			return
		}

		// Case: user provides a single index as argument (e.g., 'tc list 3')
		if len(args) > 0 && !countFlagUsed {
			index, err := strconv.Atoi(args[0])
			if err != nil || index < 1 {
				fmt.Printf("tc list: Invalid index: %s\n", args[0])
				return
			}

			items, err := manager.GetLastItems(index)
			if err != nil {
				fmt.Printf("tc list: Error retrieving clipboard history: %v\n", err)
				return
			}
			if index > len(items) {
				fmt.Printf("tc list: Index %d is out of range. Only %d items available.\n", index, len(items))
				return
			}

			fmt.Println(manager.FormatItem(items[index-1], index-1))
			return
		}

		// Default or --count: list last N items
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
	listCmd.Flags().IntP("count", "n", 10, "Number of clipboard items to list")
}
