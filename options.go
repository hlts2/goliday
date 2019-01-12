package goliday

import "time"

// Option --
type Option func(*evaluateOption)

type evaluateOption struct {
	year  int
	month int
	day   int
}

// WithYear --
func WithYear(year int) Option {
	return func(e *evaluateOption) {
		e.year = year
	}
}

// WithMonth --
func WithMonth(month int) Option {
	return func(e *evaluateOption) {
		e.month = month
	}
}

// WithDay --
func WithDay(day int) Option {
	return func(e *evaluateOption) {
		e.day = day
	}
}

// WithTime --
func WithTime(t *time.Time) Option {
	return func(e *evaluateOption) {
		e.year = t.Year()
		e.month = int(t.Month())
		e.day = t.Day()
	}
}
