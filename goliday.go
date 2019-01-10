package goliday

import (
	"fmt"
	"time"
)

// IsHolidayByTime returns true if t(time.Time) is holiday.
func IsHolidayByTime(t time.Time) bool {
	return hds.containsKey(formatYmd(t.Year(), int(t.Month()), t.Day()))
}

// IsHolidayByDate returns true if date is holiday.
func IsHolidayByDate(year, month, day int) bool {
	return hds.containsKey(formatYmd(year, month, day))
}

func formatYmd(year, month, day int) string {
	return fmt.Sprintf("%d-%d-%d", year, month, day)
}
