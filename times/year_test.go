package times

import (
	"testing"
	"time"
)

func getLocalYearRangeComparison(t time.Time) (beginning, end time.Time) {
	t = t.Local()
	tDayBg := GetLocalDayBeginning(t)

	beginning = tDayBg.AddDate(0, 0, 1-tDayBg.YearDay())
	end = beginning.AddDate(1, 0, 0)
	return
}

func TestGetLocalYearRange(t *testing.T) {
	now := time.Now()

	beginning, end := GetLocalYearRange(now)
	beginning1, end1 := getLocalYearRangeComparison(now)

	if !beginning.Equal(beginning1) || !end.Equal(end1) {
		t.Error("wrong result")
	}
}

func BenchmarkGetLocalYearRange(b *testing.B) {
	now := time.Now()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetLocalYearRange(now)
	}
}

func BenchmarkGetLocalYearRangeComparison(b *testing.B) {
	now := time.Now()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		getLocalYearRangeComparison(now)
	}
}
