package validator

import (
	"errors"
	// "fmt"
	"reflect"
	"regexp"
	"strconv"
	"time"
)

const (
	EMAIL_PATTERN = "([a-zA-Z0-9])+(@)([a-zA-Z0-9])+((.)[a-zA-Z0-9])+"
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
	if str, ok := v.Object.(string); ok {
		matched, err := regexp.MatchString(pattern, str)
		if err != nil || matched == false {
			v.Add("pattern mismatch", msg...)
		} else {
			v.Add("matching cannot be applied on this object")
		}
	} else {
		v.Add("cannot be applied on this object")
	}
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
	switch reflect.ValueOf(v.Object).Kind() {
	case reflect.Array, reflect.String:
		if max, ok := maxI.(int); ok {
			if reflect.ValueOf(v.Object).Len() > max {
				v.Add("maximum "+strconv.Itoa(max)+" characters "+"allowed", msg...)
			}
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if max, ok := maxI.(int); ok {
			if reflect.ValueOf(v.Object).Int() > int64(max) {
				v.Add("maximum "+strconv.Itoa(max)+" is allowed", msg...)
			}
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		if max, ok := maxI.(int); ok {
			if reflect.ValueOf(v.Object).Uint() > uint64(max) {
				v.Add("maximum"+strconv.Itoa(max)+"is allowed", msg...)
			}
		}
	case reflect.Float32, reflect.Float64:
		if max, ok := maxI.(int); ok {
			if reflect.ValueOf(v.Object).Float() > float64(max) {
				v.Add("maximum "+strconv.Itoa(max)+" is allowed", msg...)
			}
		}
	case reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		if max, ok := maxI.(int); ok {
			if reflect.ValueOf(v.Object).Len() > max {
				v.Add("maximum "+strconv.Itoa(max)+" numbers allowed", msg...)
			}
		}
	case reflect.Struct:
		if t, ok := maxI.(time.Time); ok {
			v.IsTimeBefore(t)
		} else {
			v.Add("cannot be applied on this object")
		}
	default:
		v.Add("cannot be applied on this object")
	}

	return v
}

func (v *Validator) Min(minI interface{}, msg ...string) *Validator {
	switch reflect.ValueOf(v.Object).Kind() {
	case reflect.Array, reflect.String:
		if min, ok := minI.(int); ok {
			if reflect.ValueOf(v.Object).Len() < min {
				v.Add("minimum "+strconv.Itoa(min)+" characters "+"required", msg...)
			}
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if min, ok := minI.(int64); ok {
			if reflect.ValueOf(v.Object).Int() < min {
				v.Add("minimum "+strconv.Itoa(int(min))+" is required", msg...)
			}
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		if min, ok := minI.(uint64); ok {
			if reflect.ValueOf(v.Object).Uint() < min {
				v.Add("minimum"+strconv.Itoa(int(min))+"is required", msg...)
			}
		}
	case reflect.Float32, reflect.Float64:
		if min, ok := minI.(float64); ok {
			if reflect.ValueOf(v.Object).Float() < min {
				v.Add("minimum "+strconv.Itoa(int(min))+" is required", msg...)
			}
		}
	case reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		if min, ok := minI.(int); ok {
			if reflect.ValueOf(v.Object).Len() < min {
				v.Add("minimum "+strconv.Itoa(min)+" numbers required", msg...)
			}
		}
	case reflect.Struct:
		if t, ok := minI.(time.Time); ok {
			v.IsTimeAfter(t)
		} else {
			v.Add("cannot be applied on this object")
		}
	default:
		v.Add("cannot be applied on this object")
	}

	return v
}

func (v *Validator) Required(msg ...string) *Validator {
	if v.Object == nil {
		v.Add("required", msg...)
	}

	// inspired by https://github.com/jamieomatthews/validation
	switch reflect.ValueOf(v.Object).Kind() {
	case reflect.Array, reflect.String:
		if reflect.ValueOf(v.Object).Len() == 0 {
			v.Add("required", msg...)
		}
	case reflect.Bool:
		if !reflect.ValueOf(v.Object).Bool() {
			v.Add("required", msg...)
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if reflect.ValueOf(v.Object).Int() == 0 {
			v.Add("required", msg...)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		if reflect.ValueOf(v.Object).Uint() == 0 {
			v.Add("required", msg...)
		}
	case reflect.Float32, reflect.Float64:
		if reflect.ValueOf(v.Object).Float() == 0 {
			v.Add("required", msg...)
		}
	case reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		if reflect.ValueOf(v.Object).IsNil() {
			v.Add("required", msg...)
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
