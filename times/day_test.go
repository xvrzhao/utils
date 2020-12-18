package times

import (
	"testing"
	"time"
)

func TestGetLocalDayRange(t *testing.T) {
	tt := time.Date(2020, 12, 18, 13, 12, 11, 0, time.Local)
	b, e := GetLocalDayRange(tt)
	if b != time.Date(2020, 12, 18, 0, 0, 0, 0, time.Local) ||
		e != time.Date(2020, 12, 19, 0, 0, 0, 0, time.Local) {
		t.Error("wrong error")
	}
}
