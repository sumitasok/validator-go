package validator

import (
	// "fmt"
	"github.com/stretchr/testify/assert"
	// "reflect"
	"testing"
	"time"
)

var (
	tArray = []string{"Lorem", "ipsum", "dolor", "sit", "amet", "consectetur", "adipiscing"}
)

// func TestMaxMixAndMatch(t *testing.T) {
// 	assert := assert.New(t)

// 	// m := Max{}

// 	// ArrayValue := []string{"7"}
// 	// StringValue := "7"
// 	IntValue := int(7)
// 	Int8Value := int8(17) // -128 to 127, Signed 8-bit integer, bytes per element 1
// 	// Int16Value := int16(7) -32,768 to 32,767, Signed 16-bit integer, bytes per element 2
// 	// Int32Value := int32(7) -2,147,483,648 to 2,147,483,647, Signed 32-bit integer, bytes per element 4
// 	// Int64Value := int64(7) -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807, Signed 64-bit intege, bytes per element 8
// 	// UintValue := uint(7)
// 	// Uint8Value := uint8(7)
// 	// Uint16Value := uint16(7)
// 	// Uint32Value := uint32(7)
// 	// Uint64Value := uint64(7)
// 	// UintptrValue := uintptr(7)
// 	// Float32Value := float32(7)
// 	// Float64Value := float64(7)
// 	// InterfaceValue := 7
// 	// MapValue := map[int]string{7}
// 	// PtrValue := 7
// 	// SliceValue := 7

// 	// m.validate(v, maxI, ...)

// 	v := false

// 	k1, ok1 := makeInterface(IntValue).(int)
// 	k1tc := int64(IntValue)
// 	fmt.Println("k1", k1, "ok1", ok1, "k1tc", k1tc)

// 	k2, ok2 := makeInterface(Int8Value).(int8)
// 	k2tc := float64(Int8Value)
// 	// k264, ok2 := makeInterface(k2tc).(float64)
// 	fmt.Println("k2", k2, "ok2", ok2, "k2tc", k2tc, float64(1.1))

// 	if k2tc > float64(makeInterface(Int8Value)) {
// 		v = true
// 	}

// 	assert.True(v)
// }

// func makeInterface(i interface{}) interface{} {
// 	return i
// }

type testStr struct{}

func TestTypeAllow(t *testing.T) {
	assert := assert.New(t)

	a := time.Now()
	b := time.Now()

	assert.True(typeAllow(a, b))

	c := int(7)

	assert.False(typeAllow(a, c))

	d := int64(7)

	assert.True(typeAllow(c, d))

	e := testStr{}

	assert.False(typeAllow(e, b))
}

func TestErrorMessages(t *testing.T) {
	assert := assert.New(t)

	v := On(tArray)

	m := Max{}
	m.validate(v, 4)

	assert.Equal("maximum 4 numbers allowed", v.Error().Error())
}
