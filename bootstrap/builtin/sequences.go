package builtin

import (
	"gitlab.com/kode4food/ale/data"
	"gitlab.com/kode4food/ale/stdlib"
)

// Seq attempts to convert the provided value to a sequence, or returns nil
func Seq(args ...data.Value) data.Value {
	if s, ok := args[0].(data.Sequence); ok && !s.IsEmpty() {
		return s
	}
	return data.Nil
}

// First returns the first value in the sequence
func First(args ...data.Value) data.Value {
	return args[0].(data.Sequence).First()
}

// Rest returns the sequence elements after the first value
func Rest(args ...data.Value) data.Value {
	return args[0].(data.Sequence).Rest()
}

// Cons prepends a value to the provided sequence
func Cons(args ...data.Value) data.Value {
	h := args[0]
	r := args[1]
	return r.(data.Sequence).Prepend(h)
}

// Conj conjoins a value to the provided sequence in some way
func Conj(args ...data.Value) data.Value {
	a0 := args[0]
	if a, ok := a0.(data.Appender); ok {
		return a.Append(args[1:]...)
	}
	s := a0.(data.Sequence)
	for _, v := range args[1:] {
		s = s.Prepend(v)
	}
	return s
}

// Append combines two sequences if the first is an Appender
func Append(args ...data.Value) data.Value {
	a := args[0].(data.Appender)
	s := args[1].(data.Sequence)
	values := stdlib.SequenceToValues(s)
	return a.Append(values...)
}

// Reverse returns a reversed copy of a Sequence
func Reverse(args ...data.Value) data.Value {
	s := args[0].(data.Sequence)
	if r, ok := s.(data.Reverser); ok {
		return r.Reverse()
	}
	var res data.Sequence = data.EmptyList
	for f, r, ok := s.Split(); ok; f, r, ok = r.Split() {
		res = res.Prepend(f)
	}
	return res
}

// Size returns the size of the provided sequence
func Size(args ...data.Value) data.Value {
	s := args[0].(data.Sequence)
	l := data.Size(s)
	return data.Integer(l)
}

// Nth returns the nth element of the provided sequence
func Nth(args ...data.Value) data.Value {
	s := args[0].(data.Indexed)
	res, _ := s.ElementAt(int(args[1].(data.Integer)))
	return res
}

// Get returns a value by key from the provided mapped sequence
func Get(args ...data.Value) data.Value {
	s := args[0].(data.Mapped)
	res, _ := s.Get(args[1])
	return res
}

// IsSeq returns whether or not the provided value is a sequence
func IsSeq(args ...data.Value) data.Value {
	if _, ok := args[0].(data.Sequence); ok {
		return data.True
	}
	return data.False
}

// IsEmpty returns whether or not the provided sequence is empty
func IsEmpty(args ...data.Value) data.Value {
	s := args[0].(data.Sequence)
	return data.Bool(s.IsEmpty())
}

// IsSized returns whether or not the provided value is a sized sequence
func IsSized(args ...data.Value) data.Value {
	_, ok := args[0].(data.SizedSequence)
	return data.Bool(ok)
}

// IsIndexed returns whether or not the provided value is an indexed sequence
func IsIndexed(args ...data.Value) data.Value {
	_, ok := args[0].(data.IndexedSequence)
	return data.Bool(ok)
}
