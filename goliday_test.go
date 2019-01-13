package goliday

import (
	"testing"
	"time"
)

func TestHolidays(t *testing.T) {
	tests := []struct {
		options  []Option
		expected holidays
	}{
		{
			options: []Option{WithYear(2050)},
			expected: holidays{
				&holiday{year: 2050, month: 1, day: 1},
				&holiday{year: 2050, month: 1, day: 10},
				&holiday{year: 2050, month: 2, day: 11},
				&holiday{year: 2050, month: 2, day: 23},
				&holiday{year: 2050, month: 3, day: 20},
				&holiday{year: 2050, month: 3, day: 21},
				&holiday{year: 2050, month: 4, day: 29},
				&holiday{year: 2050, month: 5, day: 3},
				&holiday{year: 2050, month: 5, day: 4},
				&holiday{year: 2050, month: 5, day: 5},
				&holiday{year: 2050, month: 7, day: 18},
				&holiday{year: 2050, month: 8, day: 11},
				&holiday{year: 2050, month: 9, day: 19},
				&holiday{year: 2050, month: 9, day: 23},
				&holiday{year: 2050, month: 10, day: 10},
				&holiday{year: 2050, month: 11, day: 3},
				&holiday{year: 2050, month: 11, day: 23},
			},
		},
		{
			options: []Option{WithYear(2050), WithMonth(1)},
			expected: holidays{
				&holiday{year: 2050, month: 1, day: 1},
				&holiday{year: 2050, month: 1, day: 10},
			},
		},
		{
			options: []Option{WithYear(2050), WithDay(11)},
			expected: holidays{
				&holiday{year: 2050, month: 2, day: 11},
				&holiday{year: 2050, month: 8, day: 11},
			},
		},
		{
			options: []Option{WithTime(func() time.Time {
				loc, _ := time.LoadLocation("Asia/Tokyo")
				return time.Date(2050, 11, 23, 8, 4, 18, 0, loc)
			}())},
			expected: holidays{
				&holiday{year: 2050, month: 11, day: 23},
			},
		},
		{
			options:  []Option{},
			expected: holidays{},
		},
	}

	for _, test := range tests {
		holidays := Holidays(test.options...)

		if len(test.expected) != len(holidays) {
			t.Errorf("Holidays count is wrong. expected: %v, but: %v", len(test.expected), len(holidays))
		}

		for i := 0; i < len(test.expected); i++ {
			expected := test.expected[i]
			holiday := holidays[i]

			if expected.Year() != holiday.Year() {
				t.Errorf("Holidays[%d] Year is wrong. expected: %v, but: %v", i, expected.Year(), holiday.Year())
			}

			if expected.Month() != holiday.Month() {
				t.Errorf("Holidays[%d] Month is wrong. expected: %v, but: %v", i, expected.Month(), holiday.Month())
			}

			if expected.Day() != holiday.Day() {
				t.Errorf("Holidays[%d] Day is wrong. expected: %v, but: %v", i, expected.Day(), holiday.Day())
			}
		}
	}
}
