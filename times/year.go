package times

import "time"

// GetLocalYearRange 获取 t 在本地的当年时间范围 [beginning, end)，
// beginning 和 end 均为本地时区。
func GetLocalYearRange(t time.Time) (beginning, end time.Time) {
	t = t.Local()
	beginning = time.Date(t.Year(), 1, 1, 0, 0, 0, 0, time.Local)
	end = beginning.AddDate(1, 0, 0)
	return
}
