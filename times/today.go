package times

import "time"

// GetLocalTodayBeginningTimestamp 获取当地今日开始的时间戳。
func GetLocalTodayBeginningTimestamp(now time.Time) int64 {
	now = now.Local()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).Unix()
}
