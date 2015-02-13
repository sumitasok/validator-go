package validator

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type tObj struct {
	Name string
}

func TestRequired(t *testing.T) {
	assert := assert.New(t)

	v := On("Sumit").Required()
	assert.NoError(v.Error())

	ts := tObj{}
	v1 := On(ts.Name).Required()
	assert.Error(v1.Error())
}

func TestMin(t *testing.T) {
	assert := assert.New(t)

	v2 := On("Sumit").Required().Min(7)
	assert.Equal("minimum 7 characters required",
		v2.Error().Error())
}

func TestMax(t *testing.T) {
	assert := assert.New(t)

	// replicate errors for more types

	v3int := On(123).Required().Max(7)
	assert.Equal("maximum 7 is allowed",
		v3int.Error().Error())

	v3 := On("classified").Required().Max(7)
	assert.Equal("maximum 7 characters allowed",
		v3.Error().Error())
}

func TestRange(t *testing.T) {
	assert := assert.New(t)

	v4max := On("classified").Required().Range(3, 7)
	assert.Equal("maximum 7 characters allowed",
		v4max.Error().Error())
	v4min := On("classified").Required().Range(13, 17)
	assert.Equal("minimum 13 characters required",
		v4min.Error().Error())
	v4maxInt := On(123).Required().Range(3, 7)
	assert.Equal("maximum 7 is allowed",
		v4maxInt.Error().Error())

	v5 := On(tObj{}).Required().Range(13, 17)
	assert.Equal("cannot be applied on this object",
		v5.Error().Error())
}

func TestMinTime(t *testing.T) {
	assert := assert.New(t)

	taf := time.Date(2011, time.November, 10, 23, 0, 0, 0, time.UTC)
	tbef := time.Date(2010, time.November, 10, 23, 0, 0, 0, time.UTC)

	v2tbef := On(tbef).Required().Min(taf)
	assert.Equal("time should be after 2011-11-10 23:00:00 +0000 UTC",
		v2tbef.Error().Error())

	v7 := On(tbef).IsTimeAfter(taf)
	assert.Equal("time should be after 2011-11-10 23:00:00 +0000 UTC",
		v7.Error().Error())
}

func TestMaxTime(t *testing.T) {
	assert := assert.New(t)

	taf := time.Date(2011, time.November, 10, 23, 0, 0, 0, time.UTC)
	tbef := time.Date(2010, time.November, 10, 23, 0, 0, 0, time.UTC)

	v2tbef := On(taf).Required().Max(tbef)
	assert.Equal("time should be before 2010-11-10 23:00:00 +0000 UTC",
		v2tbef.Error().Error())

	v6 := On(taf).IsTimeBefore(tbef)
	assert.Equal("time should be before 2010-11-10 23:00:00 +0000 UTC",
		v6.Error().Error())
}

func TestMatch(t *testing.T) {
	assert := assert.New(t)

	vMatch := On("obj").Required().Match("([a-zA-Z0-9])+(@)([a-zA-Z0-9])+((.)[a-zA-Z0-9])+")
	assert.Equal("pattern mismatch", vMatch.Error().Error())

	vEmail := On("obj").Required().Email()
	assert.Equal("not a valid email", vEmail.Error().Error())
}

func TestAdd(t *testing.T) {
	assert := assert.New(t)

	tO := tObj{}

	v := On(tO.Name).Key("name")
	v.Add("required field")
	assert.Equal("name: required field", v.Error().Error())
}

func TestKey(t *testing.T) {
	assert := assert.New(t)

	tO := tObj{}

	v := On(tO.Name).Key("name")
	assert.Equal("name", v.KeyName)
}

func TestError(t *testing.T) {
	assert := assert.New(t)
	v := Validator{}
	assert.Equal(nil, v.Error())
}
