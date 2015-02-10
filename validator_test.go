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

	v2 := On("Sumit").Required().Min(7)
	assert.Equal("minimum 7 characters required", v2.Error().Error())
	// replicate errors for more types

	v3 := On("classified").Required().Max(7)
	assert.Equal("maximum 7 characters allowed", v3.Error().Error())

}
