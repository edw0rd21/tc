package cmd

import (
	"fmt"

	"github.com/edw0rd21/tc/internal/clipboard"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List clipboard history items",
	Long:  `Displays the most recent clipboard history items.`,
	Run: func(cmd *cobra.Command, args []string) {
		count, _ := cmd.Flags().GetInt("count")

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
	listCmd.Flags().IntP("count", "n", 10, "Number of clipboard items to list")
}
