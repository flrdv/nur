// Code generated by "stringer -type=Operation"; DO NOT EDIT.

package operation

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Unknown-0]
	_ = x[Add-1]
	_ = x[Sub-2]
	_ = x[Mul-3]
	_ = x[Div-4]
	_ = x[Pow-5]
	_ = x[Negate-6]
	_ = x[LogicNegate-7]
}

const _Operation_name = "UnknownAddSubMulDivPowNegateLogicNegate"

var _Operation_index = [...]uint8{0, 7, 10, 13, 16, 19, 22, 28, 39}

func (i Operation) String() string {
	if i < 0 || i >= Operation(len(_Operation_index)-1) {
		return "Operation(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Operation_name[_Operation_index[i]:_Operation_index[i+1]]
}
