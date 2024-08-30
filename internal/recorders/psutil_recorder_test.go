package recorders

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PsUtilRecorder_Record_Returns_NonEmpty(t *testing.T) {
	pr := PsUtilRecorder{}

	activity, err := pr.Record()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if activity == nil {
		t.Error("Expected activity to be set")
		return
	}

	if activity.Time.IsZero() {
		t.Error("Expected Time to be set")
	}
	assert.NotEmpty(t, activity.String())
}
