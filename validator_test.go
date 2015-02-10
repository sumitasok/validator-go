package validator

import (
	// "fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidator(t *testing.T) {
	assert := assert.New(t)

	v := Validator{}
	obj := v.On("Sumit")

	assert.IsType(VString{}, v.Object)

	obj.Required()

	assert.NoError(obj.Error)

	obj = v.On("")

	assert.IsType(VString{}, v.Object)

	obj.Required()

	assert.Error(obj.Error)
}
