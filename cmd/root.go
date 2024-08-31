package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var logLevel int32

var rootCommand = &cobra.Command{
	Use:   "net-monitor",
	Short: "A simple CLI tool to monitor network interfaces activity",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use 'net-monitor monitor' to start monitoring network interfaces")
	},
}

func Execute() {
	log.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
		Level(zerolog.Level(logLevel)).
		With().
		Timestamp().
		Logger()

	if err := rootCommand.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCommand.PersistentFlags().Int32VarP(&logLevel, "log-level", "l", -1, "Log levels.\n-1: Trace\n0: Debug\n1: Info\n2: Warn\n3: Error\n4: Fatal\n5: Panic")
	rootCommand.AddCommand(monitorCmd)

}
