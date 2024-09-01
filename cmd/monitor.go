package cmd

import (
	"net"
	validator "net-monitor/internal/monitor/nic"

	"github.com/rs/zerolog/log"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Monitor network interfaces",
	Run:   showNetworkInterfaces,
}

func showNetworkInterfaces(cmd *cobra.Command, args []string) {
	log.Info().Msg("Fetching network interfaces...")
	ifaces, err := net.Interfaces()
	if err != nil {
		log.Fatal().Msgf("Error fetching network interfaces: %v", err)
	}

	prompt := promptui.Select{
		Label:     "Select network interface to monitor",
		Size:      10,
		Items:     ifaces,
		IsVimMode: true,
	}
	idx, str, err := prompt.Run()
	if err != nil {
		log.Fatal().Msgf("Error selecting network interface: %v", err)
	}
	log.Debug().Msgf("Selected interface: %d %s", idx, str)
	iface := &ifaces[idx]

	validator.EnsureSelectedInterfaceEnabled(iface)
	log.Debug().Msgf("Monitoring interface: %d %s", idx, str)
}
