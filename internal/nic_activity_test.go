package internal

import (
	"testing"
)

func TestNewNicActivity(t *testing.T) {
	na := newNicActivity()

	if na.pipe == nil {
		t.Error("Expected pipe to be initialized")
	}
}

func TestNicActivity_Accept(t *testing.T) {
	na := newNicActivity()

	sent := uint64(100)
	received := uint64(200)

	go func() {
		for activity := range na.pipe {
			if activity.BytesSent != sent {
				t.Errorf("Expected BytesSent to be %d, got %d", sent, activity.BytesSent)
			}
			if activity.BytesReceived != received {
				t.Errorf("Expected BytesReceived to be %d, got %d", received, activity.BytesReceived)
			}
		}
	}()

	na.Recorder.Record(sent, received)
}

func TestNicActivity_Report(t *testing.T) {
	na := newNicActivity()

	sent := uint64(100)
	received := uint64(200)

	go func() {
		activity := <-na.Report()

		if activity.BytesSent != sent {
			t.Errorf("Expected BytesSent to be %d, got %d", sent, activity.BytesSent)
		}
		if activity.BytesReceived != received {
			t.Errorf("Expected BytesReceived to be %d, got %d", received, activity.BytesReceived)
		}
	}()

	na.Recorder.Record(sent, received)
}
