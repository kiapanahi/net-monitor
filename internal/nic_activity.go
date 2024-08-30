package internal

import (
	models "net-monitor/internal/activity_instance"
	recorders "net-monitor/internal/recorders"
)

type ActivityReporter interface {
	Report() <-chan models.ActivityInstance
}

type NicActivity struct {
	Recorder recorders.ActivityRecorder
	Reporter ActivityReporter
	pipe     chan models.ActivityInstance
}

func (na NicActivity) Record() {
	ai, err := na.Recorder.Record()
	if err != nil {
		return
	}
	na.pipe <- *ai
}

func (na NicActivity) Report() <-chan models.ActivityInstance {
	return na.pipe
}

func newNicActivity(recorder recorders.ActivityRecorder, reporter ActivityReporter) NicActivity {
	return NicActivity{
		Recorder: recorder,
		Reporter: reporter,
		pipe:     make(chan models.ActivityInstance),
	}
}
