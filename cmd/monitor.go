package cmd

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/manifoldco/promptui"
	psutilsnet "github.com/shirou/gopsutil/net"
	"github.com/spf13/cobra"
)

var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Monitor network interfaces",
	Run:   showNetworkInterfaces,
}

func showNetworkInterfaces(cmd *cobra.Command, args []string) {
	ifaces, err := net.Interfaces()
	if err != nil {
		log.Fatalf("Error fetching network interfaces: %v", err)
	}

	prompt := promptui.Select{
		Label:     "Select network interface to monitor",
		Size:      10,
		Items:     ifaces,
		IsVimMode: true,
	}
	idx, str, err := prompt.Run()
	if err != nil {
		log.Fatalf("Error selecting network interface: %v", err)
	}
	log.Default().Printf("Selected interface: %d %s", idx, str)
	iface := &ifaces[idx]

	ensureSelectedInterfaceEnabled(iface)
	log.Default().Printf("Selected interface: %d %s", idx, str)
}

func ensureSelectedInterfaceEnabled(iface *net.Interface) {

	log.Default().Printf("ensuring interface %s is not a Loopback device, Up and Running", iface.Name)

	if (*iface).Flags&net.FlagUp == 0 {
		log.Default().Fatalf("Interface %s is down", iface.Name)
	}

	if (*iface).Flags&net.FlagLoopback == net.FlagLoopback {
		log.Default().Fatalf("Interface %s is a loopback interface", iface.Name)
	}

	if (*iface).Flags&net.FlagRunning == 0 {
		log.Default().Fatalf("Interface %s is not running", iface.Name)
	}
}

// OLD

func monitorNetworkInterfaces(cmd *cobra.Command, args []string) {
	nics, err := psutilsnet.Interfaces()

	if err != nil {
		log.Fatalf("Error fetching network interfaces: %v", err)
	}

	nicNames := make([]string, len(nics))
	for _, nic := range nics {
		nicNames = append(nicNames, nic.Name)
	}

	p := promptui.Select{
		Label: "Select network interface to monitor",
		Items: nicNames,
	}

	ifIndex, ifName, err := p.Run()

	if err != nil {
		log.Fatalf("Error selecting network interface: %v", err)
	}

	fmt.Println("Monitoring interface: ", ifIndex, ifName)

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	previousStats, err := psutilsnet.IOCounters(true)
	if err != nil {
		log.Fatalf("Error fetching initial stats: %v", err)
	}

	previousData := getInterfaceStats(ifName, previousStats)

	for range ticker.C {
		currentStats, err := psutilsnet.IOCounters(true)
		if err != nil {
			log.Fatalf("Error fetching network statistics: %v", err)
		}

		currentData := getInterfaceStats(ifName, currentStats)

		if currentData != nil && previousData != nil {
			sentBytes := currentData.BytesSent - previousData.BytesSent
			recvBytes := currentData.BytesRecv - previousData.BytesRecv

			fmt.Printf("Interface: %s | Sent: %v bytes/s | Received: %v bytes/s\n",
				ifName, sentBytes/uint64(1), recvBytes/uint64(1))
		}

		previousData = currentData
	}
}

func getInterfaceStats(ifaceName string, stats []psutilsnet.IOCountersStat) *psutilsnet.IOCountersStat {
	for _, iface := range stats {
		if iface.Name == ifaceName {
			return &iface
		}
	}
	return nil
}
