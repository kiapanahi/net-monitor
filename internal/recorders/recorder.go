package recorders

import models "net-monitor/internal/activity_instance"

type ActivityRecorder interface {
	Record() (*models.ActivityInstance, error)
}
