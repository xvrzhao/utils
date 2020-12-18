package times

import (
	"testing"
	"time"
)

func TestGetLocalMonthRange(t *testing.T) {
	now := time.Now()
	if now.Hour()-now.UTC().Hour() == 8 {
		// 东八区
		t1 := time.Date(2020, 11, 30, 23, 59, 59, 0, time.Local)
		t2 := time.Date(2020, 12, 1, 0, 0, 0, 0, time.Local)

		b, e := GetLocalMonthRange(t1)
		if b != time.Date(2020, 11, 1, 0, 0, 0, 0, time.Local) ||
			e != time.Date(2020, 12, 1, 0, 0, 0, 0, time.Local) {
			t.Error("wrong result")
			return
		}
		b, e = GetLocalMonthRange(t2)
		if b != time.Date(2020, 12, 1, 0, 0, 0, 0, time.Local) ||
			e != time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local) {
			t.Error("wrong result")
		}
	}
}
