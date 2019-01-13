package goliday

// IsHoliday returns true if there is a holiday set as an option.
func IsHoliday(ops ...Option) bool {
	return table.contains(evaluatedOption(ops...))
}

// Holidays returns Holiday slice that matches option values.
func Holidays(ops ...Option) []Holiday {
	return table.holidaysByEvaluateOption(evaluatedOption(ops...))
}

func evaluatedOption(ops ...Option) *evaluateOption {
	evalOp := newEvaluateOption()
	for _, op := range ops {
		op(evalOp)
	}
	return evalOp
}
