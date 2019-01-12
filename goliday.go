package goliday

import (
	"time"
)

// IsHoliday returns true if date is holiday.
func IsHoliday(year, month, day int) bool {
	return holidays.contains(formatYmd(year, month, day))
}

// IsHolidayTime returns true if t(time.Time) is holiday.
func IsHolidayTime(t time.Time) bool {
	return holidays.contains(formatYmd(t.Year(), int(t.Month()), t.Day()))
}

// InMonth --
func InMonth(year, month int) []Holiday {
	return holidays.inMonth(year, month)
}

// InYear --
func InYear(year int) []Holiday {
	return holidays.inYear(year)
}
