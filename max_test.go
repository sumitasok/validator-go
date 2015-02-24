package validator

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTypeAllow(t *testing.T) {
	assert := assert.New(t)

	op := "allow"

	for i, c := range tprobabs {
		for j, k := range c.Ops[op] {
			assert.Equal(k.A, typeAllow(c.Obj, k.V), fmt.Sprintf("%d %s %d", i, op, j))
		}
	}
}

func TestMaxErrorMessages(t *testing.T) {
	// all error messages on max condition failed
	assert := assert.New(t)

	op := "maxFail"

	m := Max{}

	for i, cond := range tprobabs {
		for j, s := range cond.Ops[op] {
			v := On(cond.Obj)
			m.validate(v, s.V)
			assert.Equal(s.E, v.Error().Error(), fmt.Sprintf("%d %s %d", i, op, j))
		}
	}
}

func TestMinSuccess(t *testing.T) {
	// successful min computation
	assert := assert.New(t)

	op := "maxFail"

	m := Min{}

	for i, cond := range tprobabs {
		for j, s := range cond.Ops[op] {
			v := On(cond.Obj)
			m.validate(v, s.V)
			if s.E == wont {
				assert.Equal(wont, v.Error().Error())
			} else {
				assert.Empty(v.Error(), fmt.Sprintf("%d %s %d", i, "minSuccess", j))
			}
		}
	}
}
