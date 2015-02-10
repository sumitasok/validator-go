package validator

import (
	"errors"
	// "fmt"
	"reflect"
	"strconv"
	"time"
)

type Validator struct {
	Object interface{}
	Errors []error
}

func (v *Validator) Range(min int, max int) *Validator {
	v.Min(min)
	v.Max(max)
	return v
}

func (v *Validator) IsTimeBefore(t time.Time) *Validator {
	if dTime, ok := v.Object.(time.Time); ok {
		if dTime.After(t) {
			v.Add("time should be before " + t.String())
		}
	} else {
		v.Add("cannot be applied on this object")
	}

	return v
}

func (v *Validator) IsTimeAfter(t time.Time) *Validator {
	if dTime, ok := v.Object.(time.Time); ok {
		if dTime.Before(t) {
			v.Add("time should be after " + t.String())
		}
	} else {
		v.Add("cannot be applied on this object")
	}

	return v
}

func (v *Validator) Max(maxI interface{}) *Validator {
	switch reflect.ValueOf(v.Object).Kind() {
	case reflect.Array, reflect.String:
		if max, ok := maxI.(int); ok {
			if reflect.ValueOf(v.Object).Len() > max {
				v.Add("maximum " + strconv.Itoa(max) + " characters " + "allowed")
			}
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if max, ok := maxI.(int); ok {
			if reflect.ValueOf(v.Object).Int() > int64(max) {
				v.Add("maximum " + strconv.Itoa(max) + " is allowed")
			}
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		if max, ok := maxI.(int); ok {
			if reflect.ValueOf(v.Object).Uint() > uint64(max) {
				v.Add("maximum" + strconv.Itoa(max) + "is allowed")
			}
		}
	case reflect.Float32, reflect.Float64:
		if max, ok := maxI.(int); ok {
			if reflect.ValueOf(v.Object).Float() > float64(max) {
				v.Add("maximum " + strconv.Itoa(max) + " is allowed")
			}
		}
	case reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		if max, ok := maxI.(int); ok {
			if reflect.ValueOf(v.Object).Len() > max {
				v.Add("maximum " + strconv.Itoa(max) + " numbers allowed")
			}
		}
	case reflect.Struct:
		if dTime, ok := v.Object.(time.Time); ok {
			if t, ok := maxI.(time.Time); ok {
				if dTime.After(t) {
					v.Add("time should be before " + t.String())
				}
			} else {
				v.Add("cannot be applied on this object")
			}
		} else {
			v.Add("cannot be applied on this object")
		}
	default:
		v.Add("cannot be applied on this object")
	}

	return v
}

func (v *Validator) Min(minI interface{}) *Validator {
	switch reflect.ValueOf(v.Object).Kind() {
	case reflect.Array, reflect.String:
		if min, ok := minI.(int); ok {
			if reflect.ValueOf(v.Object).Len() < min {
				v.Add("minimum " + strconv.Itoa(min) + " characters " + "required")
			}
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if min, ok := minI.(int64); ok {
			if reflect.ValueOf(v.Object).Int() < min {
				v.Add("minimum " + strconv.Itoa(int(min)) + " is required")
			}
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		if min, ok := minI.(uint64); ok {
			if reflect.ValueOf(v.Object).Uint() < min {
				v.Add("minimum" + strconv.Itoa(int(min)) + "is required")
			}
		}
	case reflect.Float32, reflect.Float64:
		if min, ok := minI.(float64); ok {
			if reflect.ValueOf(v.Object).Float() < min {
				v.Add("minimum " + strconv.Itoa(int(min)) + " is required")
			}
		}
	case reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		if min, ok := minI.(int); ok {
			if reflect.ValueOf(v.Object).Len() < min {
				v.Add("minimum " + strconv.Itoa(min) + " numbers required")
			}
		}
	case reflect.Struct:
		if dTime, ok := v.Object.(time.Time); ok {
			if t, ok := minI.(time.Time); ok {
				if dTime.Before(t) {
					v.Add("time should be after " + t.String())
				}
			} else {
				v.Add("cannot be applied on this object")
			}
		} else {
			v.Add("cannot be applied on this object")
		}
	default:
		v.Add("cannot be applied on this object")
	}

	return v
}

func (v *Validator) Required() *Validator {
	if v.Object == nil {
		v.Add("required")
	}

	// inspired by https://github.com/jamieomatthews/validation
	switch reflect.ValueOf(v.Object).Kind() {
	case reflect.Array, reflect.String:
		if reflect.ValueOf(v.Object).Len() == 0 {
			v.Add("required")
		}
	case reflect.Bool:
		if !reflect.ValueOf(v.Object).Bool() {
			v.Add("required")
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if reflect.ValueOf(v.Object).Int() == 0 {
			v.Add("required")
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		if reflect.ValueOf(v.Object).Uint() == 0 {
			v.Add("required")
		}
	case reflect.Float32, reflect.Float64:
		if reflect.ValueOf(v.Object).Float() == 0 {
			v.Add("required")
		}
	case reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		if reflect.ValueOf(v.Object).IsNil() {
			v.Add("required")
		}
	case reflect.Struct:
		if _, ok := v.Object.(time.Time); ok {
		} else {
			v.Add("cannot be applied on this object")
		}
	default:
		v.Add("cannot be applied on this object")
	}
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

func (v *Validator) Add(msg string) {
	v.Errors = append(v.Errors, errors.New(msg))
}
