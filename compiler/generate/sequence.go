package generate

import (
	"fmt"

	"github.com/kode4food/ale/compiler/encoder"
	"github.com/kode4food/ale/data"
	"github.com/kode4food/ale/env"
	"github.com/kode4food/ale/runtime/isa"
)

// Error messages
const (
	errCannotCompile = "sequence cannot be compiled: %s"
)

var (
	vectorSym = env.RootSymbol("vector")
	objectSym = env.RootSymbol("object")
	dictSym = env.RootSymbol("dict")
)

// Block encodes a set of expressions, returning only the final evaluation
func Block(e encoder.Encoder, s data.Sequence) {
	f, r, ok := s.Split()
	if !ok {
		Nil(e)
		return
	}
	Value(e, f)
	for f, r, ok = r.Split(); ok; f, r, ok = r.Split() {
		e.Emit(isa.Pop)
		Value(e, f)
	}
}

// Sequence encodes a sequence
func Sequence(e encoder.Encoder, s data.Sequence) {
	switch typed := s.(type) {
	case data.String:
		Literal(e, typed)
	case data.List:
		Call(e, typed)
	case data.Vector:
		Vector(e, typed)
	case data.Object:
		Object(e, typed)
	case data.Dict:
		Dict(e, typed)
	default:
		panic(fmt.Errorf(errCannotCompile, s))
	}
}


// Vector encodes a vector
func Vector(e encoder.Encoder, v data.Vector) {
	f := resolveBuiltIn(e, vectorSym)
	callApplicative(e, f.Call(), data.Values(v))
}

// Object encodes an object
func Object(e encoder.Encoder, a data.Object) {
	args := data.Values{}
	for f, r, ok := a.Split(); ok; f, r, ok = r.Split() {
		v := f.(data.Pair)
		args = append(args, v.Car(), v.Cdr())
	}
	f := resolveBuiltIn(e, objectSym)
	callApplicative(e, f.Call(), args)
}

func Dict(e encoder.Encoder, a data.Dict) {
	args := data.Values{}
	for f, r, ok := a.Split(); ok; f, r, ok = r.Split() {
		v := f.(data.Pair)
		args = append(args, v.Car(), v.Cdr())
	}
	f := resolveBuiltIn(e, dictSym)
	callApplicative(e, f.Call(), args)
}
