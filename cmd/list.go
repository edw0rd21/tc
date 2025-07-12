package cmd

import (
	"fmt"
	"strconv"

	"github.com/edw0rd21/tc/internal/clipboard"
	"github.com/spf13/cobra"
)

var (
	countFlag int
	fullFlag  bool
	rawFlag   bool
)

var listCmd = &cobra.Command{
	Use:   "list [index]",
	Short: "List clipboard history items",
	Long: `Displays clipboard history items.

- No args: shows last N items (default 10), truncated
- [index]: shows item at that index, truncated (or full if --full is set)
- -n / --count: number of recent items to list
- -f / --full: show full content (no truncation)
- -r / --raw: show raw content without formatting`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		manager, err := clipboard.NewManager()
		if err != nil {
			fmt.Printf("tc list: Error initializing clipboard manager: %v\n", err)
			return
		}

		// If positional index is given (e.g., `tc list 3`)
		if len(args) > 0 {
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

			item := items[index-1]

			switch {
			case rawFlag:
				fmt.Println(item.Content)
			case fullFlag:
				fmt.Printf("%d➤ [%s] %s\n", index, item.Timestamp.Format("15:04:05"), item.Content)
			default:
				fmt.Println(manager.FormatItem(item, index-1))
			}
			return
		}

		// Default listing
		items, err := manager.GetLastItems(countFlag)
		if err != nil {
			fmt.Printf("tc list: Error retrieving clipboard history: %v\n", err)
			return
		}
		if len(items) == 0 {
			fmt.Println("No clipboard history found.")
			return
		}

		for i, item := range items {
			switch {
			case rawFlag:
				fmt.Println(item.Content)
			case fullFlag:
				fmt.Printf("%d➤ [%s] %s\n", i+1, item.Timestamp.Format("15:04:05"), item.Content)
			default:
				fmt.Println(manager.FormatItem(item, i))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().IntVarP(&countFlag, "count", "n", 10, "Number of items to list")
	listCmd.Flags().BoolVarP(&fullFlag, "full", "f", false, "Show full content (no truncation)")
	listCmd.Flags().BoolVarP(&rawFlag, "raw", "r", false, "Show raw content only (no formatting)")
}
