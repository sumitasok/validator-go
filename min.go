package validator

import (
	"reflect"
	"strconv"
	"time"
)

type Min struct {
	V *Validator
}

func (m *Min) validate(v *Validator, minI interface{}, msg ...string) {
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
}
