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
	
	- --count / -n: how many recent items to show (default 10).
	- --full / -f: disables truncation.
	- --raw / -r: prints only clipboard content, no formatting.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		manager, err := clipboard.NewManager()
		if err != nil {
			fmt.Printf("tc list: Error initializing clipboard manager: %v\n", err)
			return
		}

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
			fmt.Println(manager.FormatItem(item, clipboard.FormatOptions{
				Index: index - 1,
				Full:  fullFlag,
				Raw:   rawFlag,
			}))
			return
		}

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
			fmt.Println(manager.FormatItem(item, clipboard.FormatOptions{
				Index: i,
				Full:  fullFlag,
				Raw:   rawFlag,
			}))
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().IntVarP(&countFlag, "count", "n", 10, "Number of clipboard items to list")
	listCmd.Flags().BoolVarP(&fullFlag, "full", "f", false, "Show full content (no truncation)")
	listCmd.Flags().BoolVarP(&rawFlag, "raw", "r", false, "Show only raw clipboard content (no formatting)")
}
