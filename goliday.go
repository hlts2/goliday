package goliday

import (
	"fmt"
	"time"
)

// IsHoliday returns true if date is holiday.
func IsHoliday(year, month, day int) bool {
	return hds.containsKey(formatYmd(year, month, day))
}

// IsHolidayTime returns true if t(time.Time) is holiday.
func IsHolidayTime(t time.Time) bool {
	return hds.containsKey(formatYmd(t.Year(), int(t.Month()), t.Day()))
}

func formatYmd(year, month, day int) string {
	return fmt.Sprintf("%d-%d-%d", year, month, day)
}

// InMonth --
func InMonth(year, month int) {}

// InYear --
func InYear(year int) {}
