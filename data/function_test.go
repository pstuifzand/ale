package data_test

import (
	"fmt"
	"testing"

	"gitlab.com/kode4food/ale/compiler/arity"
	"gitlab.com/kode4food/ale/data"
	"gitlab.com/kode4food/ale/internal/assert"
	. "gitlab.com/kode4food/ale/internal/assert/helpers"
)

func TestApplicativeFunction(t *testing.T) {
	as := assert.New(t)

	f1 := data.MakeApplicative(func(_ ...data.Value) data.Value {
		return S("hello!")
	}, nil)

	as.True(data.IsApplicative(f1))
	as.False(data.IsNormal(f1))
	as.Contains(":type Applicative", f1)

	as.Nil(f1.CheckArity(99))
}

func TestNormalFunction(t *testing.T) {
	as := assert.New(t)

	f1 := data.MakeNormal(func(_ ...data.Value) data.Value {
		return S("hello!")
	}, arity.MakeFixedChecker(0))

	as.True(data.IsNormal(f1))
	as.False(data.IsApplicative(f1))
	as.Contains(":type Normal", f1)

	as.Nil(f1.CheckArity(0))
	err := f1.CheckArity(2)
	as.EqualError(err, fmt.Sprintf(arity.BadFixedArity, 2, 0))
}
