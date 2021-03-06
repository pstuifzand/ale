package builtin

import (
	"github.com/kode4food/ale/data"
	"github.com/kode4food/ale/internal/sequence"
)

func makeLazyResolver(f data.Function) sequence.LazyResolver {
	return func() (data.Value, data.Sequence, bool) {
		r := f.Call()
		if r != data.Nil {
			s := r.(data.Sequence)
			if sf, sr, ok := s.Split(); ok {
				return sf, sr, true
			}
		}
		return data.Nil, data.EmptyList, false
	}
}

// LazySequence treats a function as a lazy sequence
var LazySequence = data.Applicative(func(args ...data.Value) data.Value {
	fn := args[0].(data.Function)
	resolver := makeLazyResolver(fn)
	return sequence.NewLazy(resolver)
}, 1)
