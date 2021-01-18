package builtin

import (
	"errors"

	"github.com/kode4food/ale/data"
)

// Error messages
const (
	errIndexOutOfBounds = "index out of bounds"
)

// First returns the first value in the sequence
func First(args ...data.Value) data.Value {
	if c, ok := args[0].(data.Cons); ok {
		return c.Car()
	}
	return args[0].(data.Sequence).First()
}

// Rest returns the sequence elements after the first value
func Rest(args ...data.Value) data.Value {
	if c, ok := args[0].(data.Cons); ok {
		return c.Cdr()
	}
	return args[0].(data.Sequence).Rest()
}

// Append adds a value to the end of the provided Appender
func Append(args ...data.Value) data.Value {
	a := args[0].(data.Appender)
	s := args[1]
	return a.Append(s)
}

// Reverse returns a reversed copy of a Sequence
func Reverse(args ...data.Value) data.Value {
	r := args[0].(data.Reverser)
	return r.Reverse()
}

// Length returns the element count of the provided sequence
func Length(args ...data.Value) data.Value {
	s := args[0].(data.CountedSequence)
	l := s.Count()
	return data.Integer(l)
}

// Nth returns the nth element of the provided sequence or a default
func Nth(args ...data.Value) data.Value {
	s := args[0].(data.Indexed)
	i := int(args[1].(data.Integer))
	if res, ok := s.ElementAt(i); ok {
		return res
	}
	if len(args) > 2 {
		return args[2]
	}
	panic(errors.New(errIndexOutOfBounds))
}

// Get returns a value by key from the provided mapped sequence
func Get(args ...data.Value) data.Value {
	s := args[0].(data.Mapped)
	res, _ := s.Get(args[1])
	return res
}

// IsSeq returns whether the provided value is a sequence
func IsSeq(args ...data.Value) data.Value {
	_, ok := args[0].(data.Sequence)
	return data.Bool(ok)
}

// IsEmpty returns whether the provided sequence is empty
func IsEmpty(args ...data.Value) data.Value {
	s := args[0].(data.Sequence)
	return data.Bool(s.IsEmpty())
}

// IsCounted returns whether the provided value is a counted sequence
func IsCounted(args ...data.Value) data.Value {
	_, ok := args[0].(data.CountedSequence)
	return data.Bool(ok)
}

// IsIndexed returns whether the provided value is an indexed sequence
func IsIndexed(args ...data.Value) data.Value {
	_, ok := args[0].(data.IndexedSequence)
	return data.Bool(ok)
}

// IsReverser returns whether the value is a reversible sequence
func IsReverser(args ...data.Value) data.Value {
	_, ok := args[0].(data.Reverser)
	return data.Bool(ok)
}
