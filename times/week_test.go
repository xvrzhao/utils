package times

import (
	"testing"
	"time"
)

func TestGetLocalWeekRange(t *testing.T) {
	now := time.Now()
	if now.Hour()-now.UTC().Hour() == 8 {
		mockNow := time.Date(2020, 12, 18, 16, 02, 03, 04, time.Local)
		bg, end := GetLocalWeekRange(mockNow)
		if bg != time.Date(2020, 12, 14, 0, 0, 0, 0, time.Local) ||
			end != time.Date(2020, 12, 21, 0, 0, 0, 0, time.Local) {
			t.Error("wrong result")
		}
	}
}
