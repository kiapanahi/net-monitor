package recorders

import (
	models "net-monitor/internal/activity_instance"
	"time"

	psutilsnet "github.com/shirou/gopsutil/net"
)

type PsUtilRecorder struct {
}

func (pr PsUtilRecorder) Record() (*models.ActivityInstance, error) {
	stats, err := psutilsnet.IOCounters(false)

	if err != nil {
		return nil, err
	}

	return &models.ActivityInstance{
		BytesSent:     stats[0].BytesSent,
		BytesReceived: stats[0].BytesRecv,
		Time:          time.Now(),
	}, nil
}
