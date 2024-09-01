package validator

import (
	"net"

	"github.com/rs/zerolog/log"
)

func EnsureSelectedInterfaceEnabled(iface *net.Interface) {

	log.Debug().Msgf("ensuring interface %s is not a Loopback device, Up and Running", iface.Name)

	if iface.Flags&net.FlagUp == 0 {
		log.Fatal().Msgf("Interface %s is down", iface.Name)
	}

	if iface.Flags&net.FlagLoopback == net.FlagLoopback {
		log.Fatal().Msgf("Interface %s is a loopback interface", iface.Name)
	}

	if iface.Flags&net.FlagRunning == 0 {
		log.Fatal().Msgf("Interface %s is not running", iface.Name)

	}
}
