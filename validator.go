package validator

import (
	"errors"
	// "fmt"
	"reflect"
	"time"
)

type Validator struct {
	Object  interface{}
	KeyName string
	Errors  []error
}

func (v *Validator) Range(min int, max int, msg ...string) *Validator {
	v.Min(min, msg...)
	v.Max(max, msg...)
	return v
}

func (v *Validator) IsTimeBefore(t time.Time, msg ...string) *Validator {
	if dTime, ok := v.Object.(time.Time); ok {
		if dTime.After(t) {
			v.Add("time should be before "+t.String(), msg...)
		}
	} else {
		v.Add("cannot be applied on this object", msg...)
	}

	return v
}

func (v *Validator) IsTimeAfter(t time.Time, msg ...string) *Validator {
	if dTime, ok := v.Object.(time.Time); ok {
		if dTime.Before(t) {
			v.Add("time should be after "+t.String(), msg...)
		}
	} else {
		v.Add("cannot be applied on this object", msg...)
	}

	return v
}

func (v *Validator) Match(pattern string, msg ...string) *Validator {
	match := Match{}
	match.validate(v, pattern, msg...)

	return v
}

func (v *Validator) Email(msg ...string) *Validator {
	pattern := EMAIL_PATTERN
	if len(msg) == 0 {
		msg = []string{"not a valid email"}
	}
	return v.Match(pattern, msg...)
}

func (v *Validator) Max(maxI interface{}, msg ...string) *Validator {
	max := Max{}
	max.validate(v, maxI, msg...)

	return v
}

func (v *Validator) Min(minI interface{}, msg ...string) *Validator {
	min := Min{}
	min.validate(v, minI, msg...)

	return v
}

func (v *Validator) Required(msg ...string) *Validator {
	required := Required{}
	required.validate(v, msg...)

	return v
}

func On(obj interface{}) *Validator {
	return &Validator{
		Object: obj,
	}
}

func (v Validator) Error() error {
	if len(v.Errors) == 0 {
		return nil
	}
	return v.Errors[0]
}

func (v *Validator) Add(msg string, customMsg ...string) {
	if len(customMsg) != 0 {
		msg = customMsg[0]
	}

	if v.KeyName == "" {
		v.Errors = append(v.Errors, errors.New(msg))
	} else {
		v.Errors = append(v.Errors, errors.New(v.KeyName+": "+msg))
	}
}

func (v *Validator) Key(name string) *Validator {
	v.KeyName = name
	return v
}

func (v *Validator) CheckCompatibility(obj interface{}) bool {
	if reflect.ValueOf(v.Object).Kind() != reflect.ValueOf(obj).Kind() {
		v.Add("non-evaluable comparators provided")
		return false
	} else {
		return true
	}
}
