package times

import "time"

// getLocalThisWeekDayRange 获取本周 (周一到周五) 时间范围的 [beginning, end)。
func GetLocalThisWeekDayRange(now time.Time) (beginning, end time.Time) {
	now = now.Local()
	todayBg := time.Date(now.Year(), now.Month(), now.Day(),
		0, 0, 0, 0, time.Local)

	wd := now.Weekday()
	if wd == time.Monday {
		beginning = todayBg
		end = beginning.AddDate(0, 0, 7)
		return
	}

	if wd == 0 {
		wd = 7
	}
	beginning = todayBg.AddDate(0, 0, 1-int(wd))
	end = beginning.AddDate(0, 0, 7)
	return
}
