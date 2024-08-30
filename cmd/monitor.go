package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/shirou/gopsutil/net"
	psutilsnet "github.com/shirou/gopsutil/net"
	"github.com/spf13/cobra"
)

var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Monitor network interfaces",
	Run:   monitorNetworkInterfacesUsingNetlink,
}

func monitorNetworkInterfacesUsingNetlink(cmd *cobra.Command, args []string) {
	interfaces := getNetworkInterfaces()
	if len(interfaces) == 0 {
		log.Fatalf("No network interfaces found")
	}

	// Display available network interfaces for selection
	prompt := promptui.Select{
		Label: "Select network interface to monitor",
		Items: interfaces,
	}

	index, _, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed: %v", err)
	}

	selectedInterface := interfaces[index]
	log.Printf("Monitoring network interface: %s\n", selectedInterface)

	// Start monitoring the selected interface
	monitorInterface(selectedInterface)
}

func getNetworkInterfaces() []string {
	var interfaces []string
	switch runtime.GOOS {
	case "windows":
		interfaces = getWindowsInterfaces()
	case "linux":
		interfaces = getLinuxInterfaces()
	case "darwin":
		interfaces = getDarwinInterfaces()
	default:
		log.Fatalf("Unsupported platform: %s", runtime.GOOS)
	}
	return interfaces
}

func getWindowsInterfaces() []string {
	output, err := exec.Command("ipconfig").Output()
	if err != nil {
		log.Fatalf("Error executing ipconfig: %v", err)
	}
	return parseWindowsInterfaces(string(output))
}

func getLinuxInterfaces() []string {
	output, err := exec.Command("bash", "-c", "ip -o -4 addr list | awk '{print $2}' | uniq").Output()
	if err != nil {
		log.Fatalf("Error executing ip command: %v", err)
	}
	return strings.Fields(string(output))
}

func getDarwinInterfaces() []string {
	output, err := exec.Command("ifconfig", "-l").Output()
	if err != nil {
		log.Fatalf("Error executing ifconfig: %v", err)
	}
	return strings.Fields(string(output))
}

func parseWindowsInterfaces(output string) []string {
	var interfaces []string
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		if strings.Contains(line, "adapter") {
			parts := strings.Split(line, "adapter")
			iface := strings.TrimSpace(parts[1])
			iface = strings.Trim(iface, ":")
			interfaces = append(interfaces, iface)
		}
	}
	return interfaces
}

func monitorInterface(ifaceName string) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for range ticker.C {
		switch runtime.GOOS {
		case "windows":
			monitorWindowsInterface(ifaceName)
		case "linux":
			monitorLinuxInterface(ifaceName)
		case "darwin":
			monitorDarwinInterface(ifaceName)
		}
	}
}

func monitorWindowsInterface(ifaceName string) {
	// You can extend this to parse statistics for Windows, possibly using `Get-NetAdapterStatistics`
	fmt.Printf("Monitoring interface on Windows: %s\n", ifaceName)
}

func monitorLinuxInterface(ifaceName string) {
	// Similar to Windows, parse /sys/class/net/<iface>/statistics/
	fmt.Printf("Monitoring interface on Linux: %s\n", ifaceName)
}

func monitorDarwinInterface(ifaceName string) {
	// Use `netstat -ib` or `ifconfig` to fetch interface stats
	fmt.Printf("Monitoring interface on macOS: %s\n", ifaceName)
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

	previousStats, err := net.IOCounters(true)
	if err != nil {
		log.Fatalf("Error fetching initial stats: %v", err)
	}

	previousData := getInterfaceStats(ifName, previousStats)

	for range ticker.C {
		currentStats, err := net.IOCounters(true)
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

func getInterfaceStats(ifaceName string, stats []net.IOCountersStat) *net.IOCountersStat {
	for _, iface := range stats {
		if iface.Name == ifaceName {
			return &iface
		}
	}
	return nil
}
