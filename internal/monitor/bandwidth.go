package monitor

import (
	"net"

	"github.com/rs/zerolog/log"
)

func MonitorBandwidth(iface *net.Interface) error {
	if iface == nil {
		log.Panic().Msg("Interface is nil")
	}
	l := log.With().Caller().Logger()
	l.Info().Msgf("Monitoring interface: %s", iface.Name)

	return nil
}
