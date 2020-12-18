package times

import "time"

// GetLocalWeekRange 获取 t 在本地的当周 (周一到周日) 时间范围 [beginning, end)，
// beginning 和 end 均为本地时区。
func GetLocalWeekRange(t time.Time) (beginning, end time.Time) {
	t = t.Local()
	tDayBg := GetLocalDayBeginning(t)

	wd := t.Weekday()
	if wd == time.Monday {
		beginning = tDayBg
		end = beginning.AddDate(0, 0, 7)
		return
	}

	if wd == 0 {
		wd = 7
	}
	beginning = tDayBg.AddDate(0, 0, 1-int(wd))
	end = beginning.AddDate(0, 0, 7)
	return
}
