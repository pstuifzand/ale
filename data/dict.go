package data

import (
	"bytes"
	"errors"

	hamt "github.com/pstuifzand/go-hamt/hamt64"
)

type Dict struct {
	hamt.Hamt
}

func NewHamtDict(args ...Value) Value {
	var h hamt.Hamt
	h = hamt.NewFunctional(hamt.FixedTables)
	var ok bool
	var h2 hamt.Hamt
	if len(args)%2 != 0 {
		panic(errors.New(ErrMapNotPaired))
	}
	for i := len(args) - 2; i >= 0; i -= 2 {
		arg := NewCons(args[i], args[i+1])
		h2, ok = h.Put(hamt.StringKey(args[i].String()), arg)
		if ok {
			h = h2
		}
	}
	return Value(Dict{h})
}

func (t Dict) Get(value Value) (Value, bool) {
	v, ok := t.Hamt.Get(hamt.StringKey(value.String()))
	if !ok {
		return Nil, false
	}
	return v.(Cons).Cdr(), ok
}

func (t Dict) First() Value {
	var v Value = Nil
	t.Hamt.Range(func(key hamt.KeyI, value interface{}) bool {
		v = value.(Value)
		return false
	})
	return v
}

func (t Dict) Rest() Sequence {
	var k hamt.KeyI = nil
	t.Hamt.Range(func(key hamt.KeyI, value interface{}) bool {
		k = key
		return false
	})
	if k == nil {
		return Nil
	}
	h2, _, b := t.Hamt.Del(k)
	if b {
		return Dict{h2}
	}
	return Nil
}

func (t Dict) Split() (Value, Sequence, bool) {
	var k hamt.KeyI = nil
	t.Hamt.Range(func(key hamt.KeyI, value interface{}) bool {
		k = key
		return false
	})
	if k == nil {
		return Nil, Nil, false
	}
	h2, v, b := t.Hamt.Del(k)
	if b {
		return v.(Value), Dict{h2}, true
	}
	return Nil, Nil, false
}

func (t Dict) Append(value Value) Sequence {
	if c, ok := value.(Cons); ok {
		h2, _ := t.Hamt.Put(hamt.StringKey(c.Car().String()), value)
		return Sequence(Dict{h2})
	}
	panic("expected cons")
}

func (t Dict) String() string {
	var buf bytes.Buffer
	buf.WriteString("(dict ")
	first := true
	t.Hamt.Range(func(key hamt.KeyI, value interface{}) bool {
		if !first {
			buf.WriteString(" ")
		}
		c := value.(Cons)
		buf.WriteString(c.Car().String())
		buf.WriteString(" ")
		buf.WriteString(c.Cdr().String())
		first = false
		return true
	})
	buf.WriteString(")")
	return buf.String()
}

// Call turns Dict into a callable type
func (o Dict) Call() Call {
	return makeMappedCall(o)
}

// Convention returns the function's calling convention
func (o Dict) Convention() Convention {
	return ApplicativeCall
}

// CheckArity performs a compile-time arity check for the function
func (o Dict) CheckArity(argCount int) error {
	return checkRangedArity(1, 2, argCount)
}

// Count
func (o Dict) Count() int {
	return int(o.Hamt.Nentries())
}
