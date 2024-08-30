package internal

import (
	"fmt"
	"time"
)

// ActivityInstance is a struct that represents sent and received bytes in a given time
type ActivityInstance struct {
	Time          time.Time
	BytesSent     uint64
	BytesReceived uint64
}

func (ai ActivityInstance) String() string {
	return fmt.Sprintf("Time: %v - Sent: %v - Received: %v", ai.Time, ai.BytesSent, ai.BytesReceived)
}

// NewActivityInstance creates a new ActivityInstance with the given sent and received bytes for the current time
func NewActivityInstance(sent, received uint64) ActivityInstance {
	return ActivityInstance{
		Time:          time.Now(),
		BytesSent:     sent,
		BytesReceived: received,
	}
}
