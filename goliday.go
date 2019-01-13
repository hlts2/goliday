package goliday

// IsHoliday --
func IsHoliday(ops ...Option) bool {
	evalOp := newEvaluateOption()

	for _, op := range ops {
		op(evalOp)
	}

	return holidays.contains(evalOp)
}
