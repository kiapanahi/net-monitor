package cmd

import (
	"net"

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
	log.Info().Msgf("Selected interface: %d %s", idx, str)
	iface := &ifaces[idx]

	ensureSelectedInterfaceEnabled(iface)
	log.Info().Msgf("Monitoring interface: %d %s", idx, str)
}

func ensureSelectedInterfaceEnabled(iface *net.Interface) {

	log.Info().Msgf("ensuring interface %s is not a Loopback device, Up and Running", iface.Name)

	if (*iface).Flags&net.FlagUp == 0 {
		log.Error().Msgf("Interface %s is down", iface.Name)
	}

	if (*iface).Flags&net.FlagLoopback == net.FlagLoopback {
		log.Error().Msgf("Interface %s is a loopback interface", iface.Name)
	}

	if (*iface).Flags&net.FlagRunning == 0 {
		log.Error().Msgf("Interface %s is not running", iface.Name)
	}
}
