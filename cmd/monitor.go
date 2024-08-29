package cmd

import (
	"fmt"
	"github.com/shirou/gopsutil/net"
	"github.com/spf13/cobra"
	"log"
	"time"
)

var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Monitor network interfaces",
	Run:   monitorNetworkInterfaces,
}

func monitorNetworkInterfaces(cmd *cobra.Command, args []string) {
	for {
		interfaces, err := net.IOCounters(true)
		if err != nil {
			log.Fatalf("Error fetching network statistics: %v", err)
		}

		for _, nic := range interfaces {

			fmt.Printf("Interface: %s\n", nic.Name)
			fmt.Printf("  Bytes Sent: %v\n", nic.BytesSent)
			fmt.Printf("  Bytes Received: %v\n", nic.BytesRecv)
			fmt.Println()
		}

		time.Sleep(1 * time.Second) // Wait for 1 second before refreshing
	}
}
