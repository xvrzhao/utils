package times

import (
	"testing"
	"time"
)

func TestGetLocalTodayBeginningTimestamp(t *testing.T) {
	now := time.Now()
	res0 := GetLocalTodayBeginningTimestamp(now)
	res1 := GetLocalTodayBeginningTimestamp(now.UTC())

	layout := "20060102"
	bt, err := time.ParseInLocation(layout, now.Format(layout), time.Local)
	if err != nil {
		t.Errorf("failed to parse time: %v", err)
		return
	}

	res2 := bt.Unix()
	if !(res0 == res1 && res1 == res2) {
		t.Errorf("res0: %d, res1: %d, res2: %d", res0, res1, res2)
		return
	}

	t.Log(res0)
}
