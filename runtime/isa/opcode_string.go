// Code generated by "stringer -type=Opcode"; DO NOT EDIT.

package isa

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Label-256]
	_ = x[Add-0]
	_ = x[Arg-1]
	_ = x[ArgLen-2]
	_ = x[Bind-3]
	_ = x[Call-4]
	_ = x[Call0-5]
	_ = x[Call1-6]
	_ = x[Closure-7]
	_ = x[CondJump-8]
	_ = x[Const-9]
	_ = x[Declare-10]
	_ = x[Div-11]
	_ = x[Dup-12]
	_ = x[Eq-13]
	_ = x[False-14]
	_ = x[Gt-15]
	_ = x[Gte-16]
	_ = x[Jump-17]
	_ = x[Load-18]
	_ = x[Lt-19]
	_ = x[Lte-20]
	_ = x[MakeCall-21]
	_ = x[MakeTruthy-22]
	_ = x[Mod-23]
	_ = x[Mul-24]
	_ = x[Neg-25]
	_ = x[NegInf-26]
	_ = x[NegOne-27]
	_ = x[Neq-28]
	_ = x[Nil-29]
	_ = x[NoOp-30]
	_ = x[Not-31]
	_ = x[One-32]
	_ = x[Panic-33]
	_ = x[Pop-34]
	_ = x[PosInf-35]
	_ = x[Resolve-36]
	_ = x[RestArg-37]
	_ = x[Return-38]
	_ = x[RetFalse-39]
	_ = x[RetNil-40]
	_ = x[RetTrue-41]
	_ = x[Self-42]
	_ = x[Store-43]
	_ = x[Sub-44]
	_ = x[TailCall-45]
	_ = x[True-46]
	_ = x[Two-47]
	_ = x[Zero-48]
}

const (
	_Opcode_name_0 = "AddArgArgLenBindCallCall0Call1ClosureCondJumpConstDeclareDivDupEqFalseGtGteJumpLoadLtLteMakeCallMakeTruthyModMulNegNegInfNegOneNeqNilNoOpNotOnePanicPopPosInfResolveRestArgReturnRetFalseRetNilRetTrueSelfStoreSubTailCallTrueTwoZero"
	_Opcode_name_1 = "Label"
)

var (
	_Opcode_index_0 = [...]uint8{0, 3, 6, 12, 16, 20, 25, 30, 37, 45, 50, 57, 60, 63, 65, 70, 72, 75, 79, 83, 85, 88, 96, 106, 109, 112, 115, 121, 127, 130, 133, 137, 140, 143, 148, 151, 157, 164, 171, 177, 185, 191, 198, 202, 207, 210, 218, 222, 225, 229}
)

func (i Opcode) String() string {
	switch {
	case 0 <= i && i <= 48:
		return _Opcode_name_0[_Opcode_index_0[i]:_Opcode_index_0[i+1]]
	case i == 256:
		return _Opcode_name_1
	default:
		return "Opcode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
