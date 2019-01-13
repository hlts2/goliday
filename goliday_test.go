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
			options: []Option{WithMonth(8)},
			expected: holidays{
				&holiday{year: 2016, month: 8, day: 11},
				&holiday{year: 2017, month: 8, day: 11},
				&holiday{year: 2018, month: 8, day: 11},
				&holiday{year: 2019, month: 8, day: 11},
				&holiday{year: 2019, month: 8, day: 12},
				&holiday{year: 2020, month: 8, day: 10},
				&holiday{year: 2021, month: 8, day: 11},
				&holiday{year: 2022, month: 8, day: 11},
				&holiday{year: 2023, month: 8, day: 11},
				&holiday{year: 2024, month: 8, day: 11},
				&holiday{year: 2024, month: 8, day: 12},
				&holiday{year: 2025, month: 8, day: 11},
				&holiday{year: 2026, month: 8, day: 11},
				&holiday{year: 2027, month: 8, day: 11},
				&holiday{year: 2028, month: 8, day: 11},
				&holiday{year: 2029, month: 8, day: 11},
				&holiday{year: 2030, month: 8, day: 11},
				&holiday{year: 2030, month: 8, day: 12},
				&holiday{year: 2031, month: 8, day: 11},
				&holiday{year: 2032, month: 8, day: 11},
				&holiday{year: 2033, month: 8, day: 11},
				&holiday{year: 2034, month: 8, day: 11},
				&holiday{year: 2035, month: 8, day: 11},
				&holiday{year: 2036, month: 8, day: 11},
				&holiday{year: 2037, month: 8, day: 11},
				&holiday{year: 2038, month: 8, day: 11},
				&holiday{year: 2039, month: 8, day: 11},
				&holiday{year: 2040, month: 8, day: 11},
				&holiday{year: 2041, month: 8, day: 11},
				&holiday{year: 2041, month: 8, day: 12},
				&holiday{year: 2042, month: 8, day: 11},
				&holiday{year: 2043, month: 8, day: 11},
				&holiday{year: 2044, month: 8, day: 11},
				&holiday{year: 2045, month: 8, day: 11},
				&holiday{year: 2046, month: 8, day: 11},
				&holiday{year: 2047, month: 8, day: 11},
				&holiday{year: 2047, month: 8, day: 12},
				&holiday{year: 2048, month: 8, day: 11},
				&holiday{year: 2049, month: 8, day: 11},
				&holiday{year: 2050, month: 8, day: 11},
			},
		},
		{
			options: []Option{WithDay(30)},
			expected: holidays{
				&holiday{year: 1973, month: 4, day: 30},
				&holiday{year: 1979, month: 4, day: 30},
				&holiday{year: 1984, month: 4, day: 30},
				&holiday{year: 1990, month: 4, day: 30},
				&holiday{year: 2001, month: 4, day: 30},
				&holiday{year: 2007, month: 4, day: 30},
				&holiday{year: 2012, month: 4, day: 30},
				&holiday{year: 2018, month: 4, day: 30},
				&holiday{year: 2029, month: 4, day: 30},
				&holiday{year: 2035, month: 4, day: 30},
				&holiday{year: 2040, month: 4, day: 30},
				&holiday{year: 2046, month: 4, day: 30},
			},
		},
		{
			options: []Option{WithTime(func() *time.Time {
				loc, _ := time.LoadLocation("Asia/Tokyo")
				t := time.Date(2050, 11, 23, 8, 4, 18, 0, loc)
				return &t
			}())},
			expected: holidays{
				&holiday{year: 2050, month: 11, day: 23},
			},
		},
	}

	for i, test := range tests {
		holidays := Holidays(test.options...)

		if len(test.expected) != len(holidays) {
			t.Errorf("[%d] - Holidays count is wrong. expected: %v, but: %v", i, len(test.expected), len(holidays))
		}

		for j := 0; i < len(test.expected); i++ {
			expected := test.expected[j]
			holiday := holidays[j]

			if expected.Year() != holiday.Year() {
				t.Errorf("[%d] - Holidays[%d] Year is wrong. expected: %v, but: %v", i, j, expected.Year(), holiday.Year())
			}

			if expected.Month() != holiday.Month() {
				t.Errorf("[%d] - Holidays[%d] Month is wrong. expected: %v, but: %v", i, j, expected.Month(), holiday.Month())
			}

			if expected.Day() != holiday.Day() {
				t.Errorf("[%d] - Holidays[%d] Day is wrong. expected: %v, but: %v", i, j, expected.Day(), holiday.Day())
			}
		}
	}
}

func TestIsHoliday(t *testing.T) {
	tests := []struct {
		options  []Option
		expected bool
	}{
		{
			options:  []Option{WithYear(2050)},
			expected: true,
		},
		{
			options:  []Option{WithYear(2050), WithMonth(1)},
			expected: true,
		},
		{
			options:  []Option{WithYear(2050), WithDay(11)},
			expected: true,
		},
		{
			options: []Option{WithTime(func() *time.Time {
				loc, _ := time.LoadLocation("Asia/Tokyo")
				t := time.Date(2050, 11, 23, 8, 4, 18, 0, loc)
				return &t
			}())},
			expected: true,
		},
		{
			options:  []Option{},
			expected: false,
		},
		{
			options: []Option{
				WithYear(1),
			},
			expected: false,
		},
		{
			options:  []Option{WithYear(2050), WithMonth(12)},
			expected: false,
		},
		{
			options:  []Option{WithYear(2050), WithDay(30)},
			expected: false,
		},
	}

	for _, test := range tests {
		got := IsHoliday(test.options...)

		if test.expected != got {
			t.Errorf("IsHoliday is wrong. expected: %v, got: %v", test.expected, got)
		}
	}
}
