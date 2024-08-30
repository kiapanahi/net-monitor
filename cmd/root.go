package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "net-monitor",
	Short: "A simple CLI tool to monitor network interfaces activity",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use 'net-monitor monitor' to start monitoring network interfaces")
	},
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCommand.AddCommand(monitorCmd)
}
