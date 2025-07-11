package cmd

import (
	"fmt"

	"github.com/edw0rd21/tc/internal/daemon"
	"github.com/spf13/cobra"
)

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Start clipboard watcher in background",
	Long:  `Starts a background process that watches your clipboard and stores history.`,
	Run: func(cmd *cobra.Command, args []string) {
		watcher, err := daemon.NewWatcher()
		if err != nil {
			fmt.Println("tc daemon: Error initializing watcher:", err)
			return
		}
		fmt.Println("tc daemon: Clipboard watcher started...")
		watcher.Start() // blocks forever
	},
}

func init() {
	rootCmd.AddCommand(daemonCmd)
}
