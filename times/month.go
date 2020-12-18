package times

import "time"

// GetLocalMonthRange 获取 t 在本地的当月时间范围 [beginning, end)，
// beginning 和 end 均为本地时区。
func GetLocalMonthRange(t time.Time) (beginning, end time.Time) {
	t = t.Local()
	tDayBg := GetLocalDayBeginning(t)

	beginning = tDayBg.AddDate(0, 0, 1-t.Day())
	end = beginning.AddDate(0, 1, 0)
	return
}
