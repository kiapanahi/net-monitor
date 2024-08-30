package internal

import (
	"testing"
	"time"
)

func TestNewActivityInstance(t *testing.T) {
	sent := uint64(300)
	received := uint64(400)

	ai := NewActivityInstance(sent, received)

	expectedTime := time.Now()
	if !ai.Time.Equal(expectedTime) {
		t.Errorf("Expected Time to be %v, but got %v", expectedTime, ai.Time)
	}

	if ai.BytesSent != sent {
		t.Errorf("Expected BytesSent to be %d, but got %d", sent, ai.BytesSent)
	}

	if ai.BytesReceived != received {
		t.Errorf("Expected BytesReceived to be %d, but got %d", received, ai.BytesReceived)
	}
}
