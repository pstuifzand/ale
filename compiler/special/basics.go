package special

import (
	"fmt"

	"gitlab.com/kode4food/ale/compiler/arity"
	"gitlab.com/kode4food/ale/compiler/encoder"
	"gitlab.com/kode4food/ale/compiler/generate"
	"gitlab.com/kode4food/ale/data"
	"gitlab.com/kode4food/ale/eval"
	"gitlab.com/kode4food/ale/namespace"
	"gitlab.com/kode4food/ale/runtime/isa"
)

// Error messages
const (
	UnpairedBindings = "bindings must be paired"
)

// Eval encodes an evaluation
func Eval(e encoder.Type, args ...data.Value) {
	arity.AssertFixed(1, len(args))
	generate.Value(e, args[0])
	generate.Literal(e, evalFor(e.Globals()))
	e.Emit(isa.Call1)
}

func evalFor(ns namespace.Type) data.Call {
	return data.Call(func(args ...data.Value) data.Value {
		return eval.Value(ns, args[0])
	})
}

// Do encodes a set of expressions, returning only the final evaluation
func Do(e encoder.Type, args ...data.Value) {
	generate.Block(e, data.Vector(args))
}

// If encodes an (if cond then else) form
func If(e encoder.Type, args ...data.Value) {
	al := arity.AssertRanged(2, 3, len(args))
	generate.Branch(e,
		func() {
			generate.Value(e, args[0])
			e.Emit(isa.MakeTruthy)
		},
		func() {
			generate.Value(e, args[1])
		},
		func() {
			if al == 3 {
				generate.Value(e, args[2])
			} else {
				generate.Nil(e)
			}
		},
	)
}

// Let encodes a (let [bindings] & body) form
func Let(e encoder.Type, args ...data.Value) {
	arity.AssertMinimum(2, len(args))
	bindings := args[0].(data.Vector)
	lb := len(bindings)
	if lb%2 != 0 {
		panic(fmt.Errorf(UnpairedBindings))
	}

	for i := 0; i < lb; i += 2 {
		n := bindings[i].(data.LocalSymbol).Name()
		e.PushLocals()
		generate.Value(e, bindings[i+1])
		e.Emit(isa.Store, e.AddLocal(n))
	}

	body := data.Vector(args[1:])
	generate.Block(e, body)

	for i := 0; i < lb; i += 2 {
		e.PopLocals()
	}
}