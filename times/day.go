package times

import "time"

// GetLocalDayRange 获取 t 在本地的当日时间范围 [beginning, end)，
// beginning 和 end 均为本地时区。
func GetLocalDayRange(t time.Time) (beginning, end time.Time) {
	beginning = GetLocalDayBeginning(t)
	end = beginning.AddDate(0, 0, 1)
	return
}

// GetLocalDayBeginning 获取 t 在本地的当日开始时间，
// 返回的时间为本地时区。
func GetLocalDayBeginning(t time.Time) time.Time {
	t = t.Local()
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
}
