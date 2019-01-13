package goliday

// IsHoliday --
func IsHoliday(ops ...Option) bool {
	evalOp := newEvaluateOption()
	for _, op := range ops {
		op(evalOp)
	}

	return table.contains(evalOp)
}

// Holidays --
func Holidays(ops ...Option) []Holiday {
	evalOp := newEvaluateOption()
	for _, op := range ops {
		op(evalOp)
	}

	return table.holidaysByEvaluateOption(evalOp)
}
