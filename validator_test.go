package validator

import (
	// "fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidator(t *testing.T) {
	assert := assert.New(t)

	v := Validator{}
	v.Work("Sumit")

	assert.IsType(VString{}, v.Object)
}
