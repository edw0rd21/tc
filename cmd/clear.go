package cmd

import (
	"fmt"

	"github.com/edw0rd21/tc/internal/clipboard"
	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear clipboard history",
	Long:  `Clears all stored clipboard history items from local storage.`,
	Run: func(cmd *cobra.Command, args []string) {
		manager, err := clipboard.NewManager()
		if err != nil {
			fmt.Printf("tc clear: Error initializing clipboard manager: %v\n", err)
			return
		}

		err = manager.ClearHistory()
		if err != nil {
			fmt.Printf("tc clear: Error clearing clipboard history: %v\n", err)
			return
		}

		quiet, _ := cmd.Flags().GetBool("quiet")
		if !quiet {
			fmt.Println("Clipboard history cleared successfully")
		}
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)
	clearCmd.Flags().BoolP("quiet", "q", false, "Suppress output messages")
}
