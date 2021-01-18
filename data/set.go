package data

import (
	"bytes"

	hamt "github.com/pstuifzand/go-hamt/hamt64"
)

type Set struct {
	hamt.Hamt
}

func (t Set) First() Value {
	var v Value = Nil
	t.Hamt.Range(func(key hamt.KeyI, value interface{}) bool {
		v = value.(Value)
		return false
	})
	return v
}

func (t Set) Rest() Sequence {
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
		return Set{h2}
	}
	return Nil
}

func (t Set) Split() (Value, Sequence, bool) {
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
		return v.(Value), Set{h2}, true
	}
	return Nil, Nil, false
}

func (t Set) Append(value Value) Sequence {
	h2, _ := t.Hamt.Put(hamt.StringKey(value.String()), value)
	return Sequence(Set{h2})
}

func (t Set) String() string {
	var buf bytes.Buffer
	buf.WriteString("(set")
	first := true
	t.Hamt.Range(func(key hamt.KeyI, value interface{}) bool {
		buf.WriteString(" ")
		buf.WriteString(value.(Value).String())
		first = false
		return true
	})
	buf.WriteString(")")
	return buf.String()
}

func NewHamtSet(args ...Value) Value {
	var h hamt.Hamt
	h = hamt.NewFunctional(hamt.FixedTables)
	var ok bool
	var h2 hamt.Hamt
	for _, arg := range args {
		h2, ok = h.Put(hamt.StringKey(arg.String()), arg)
		if ok {
			h = h2
		}
	}
	return Value(Set{h})
}
