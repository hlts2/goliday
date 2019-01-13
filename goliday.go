package goliday

// IsHoliday returns true if there is a holiday set as an option.
func IsHoliday(ops ...Option) bool {
	evalOp := newEvaluateOption()
	for _, op := range ops {
		op(evalOp)
	}

	return table.contains(evalOp)
}

// Holidays returns Holiday slice that matches option values.
func Holidays(ops ...Option) []Holiday {
	evalOp := newEvaluateOption()
	for _, op := range ops {
		op(evalOp)
	}

	return table.holidaysByEvaluateOption(evalOp)
}
