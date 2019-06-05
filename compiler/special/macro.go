package special

import (
	"gitlab.com/kode4food/ale/compiler/arity"
	"gitlab.com/kode4food/ale/compiler/encoder"
	"gitlab.com/kode4food/ale/compiler/generate"
	"gitlab.com/kode4food/ale/data"
	"gitlab.com/kode4food/ale/macro"
	"gitlab.com/kode4food/ale/namespace"
	"gitlab.com/kode4food/ale/runtime/isa"
	"gitlab.com/kode4food/ale/runtime/vm"
)

// Macro encodes a macro
func Macro(e encoder.Type, args ...data.Value) {
	vars := parseLambda(args)
	fe := makeLambdaEncoder(e, vars)
	fe.encodeCall()
	generate.Literal(e, data.Call(func(args ...data.Value) data.Value {
		closure := args[0].(*vm.Closure)
		body := closure.Caller()
		arityChecker := closure.ArityChecker
		wrapper := func(_ namespace.Type, args ...data.Value) data.Value {
			if err := arityChecker(len(args)); err != nil {
				panic(err)
			}
			return body(args...)
		}
		return macro.Call(wrapper)
	}))
	e.Emit(isa.Call1)
}

// Quote converts its argument into a literal value
func Quote(e encoder.Type, args ...data.Value) {
	arity.AssertFixed(1, len(args))
	generate.Literal(e, args[0])
}

// MacroExpand performs macro expansion of a form until it can no longer
func MacroExpand(e encoder.Type, args ...data.Value) {
	arity.AssertFixed(1, len(args))
	generate.Value(e, args[0])
	generate.Literal(e, expandFor(e.Globals()))
	e.Emit(isa.Call1)
}

func expandFor(ns namespace.Type) data.Call {
	return data.Call(func(args ...data.Value) data.Value {
		return macro.Expand(ns, args[0])
	})
}

// MacroExpand1 performs a single-step macro expansion of a form
func MacroExpand1(e encoder.Type, args ...data.Value) {
	arity.AssertFixed(1, len(args))
	generate.Value(e, args[0])
	generate.Literal(e, expand1For(e.Globals()))
	e.Emit(isa.Call1)
}

func expand1For(ns namespace.Type) data.Call {
	return data.Call(func(args ...data.Value) data.Value {
		return macro.Expand1(ns, args[0])
	})
}
