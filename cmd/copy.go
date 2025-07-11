package cmd

import (
	"fmt"
	"strconv"

	"github.com/edw0rd21/tc/internal/clipboard"

	"github.com/spf13/cobra"
)

var copyCmd = &cobra.Command{
	Use:   "copy <index>",
	Short: "Copy a specific item on clipboard",
	Long:  `Copy a specific item from clipboard history back to the system clipboard.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		index, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("tc copy: Invalid index: %s\n", args[0])
			return
		}

		if index < 1 {
			fmt.Println("tc copy: Index must be 1 or greater")
			return
		}

		manager, err := clipboard.NewManager()
		if err != nil {
			fmt.Printf("tc copy: Error initializing clipboard manager: %v\n", err)
			return
		}

		// Get the limit from flag or default to 100
		limit, _ := cmd.Flags().GetInt("limit")
		items, err := manager.GetLastItems(limit)
		if err != nil {
			fmt.Printf("tc copy: Error getting clipboard history: %v\n", err)
			return
		}

		if index > len(items) {
			fmt.Printf("tc copy: Index %d is out of range. Only %d items available.\n", index, len(items))
			return
		}

		// Convert to 0-based index
		item := items[index-1]

		err = manager.CopyToClipboard(item.Content)
		if err != nil {
			fmt.Printf("tc copy: Error copying to clipboard: %v\n", err)
			return
		}

		// Check if we should be quiet
		quiet, _ := cmd.Flags().GetBool("quiet")
		if !quiet {
			fmt.Printf("Copied item %d to clipboard\n", index)
		}
	},
}

func init() {
	rootCmd.AddCommand(copyCmd)
	copyCmd.Flags().IntP("limit", "l", 100, "Maximum number of items to search through")
	copyCmd.Flags().BoolP("quiet", "q", false, "Suppress output messages")
}
