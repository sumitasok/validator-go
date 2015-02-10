package validator

import (
	// "fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestStruct struct {
	Name string
}

func TestValidator(t *testing.T) {
	assert := assert.New(t)

	v := On("Sumit").Required()
	assert.NoError(v.Error())

	ts := TestStruct{}
	v1 := On(ts.Name).Required()
	assert.Error(v1.Error())
}
