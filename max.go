package validator

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

type Max struct {
	V *Validator
}

func (m *Max) Allow(obj interface{}) bool {
	objType := reflect.ValueOf(obj).Kind()
	valueType := reflect.ValueOf(m.V.Object).Kind()
	fmt.Println(objType)
	switch objType {
	case reflect.Int:
		switch valueType {
		case reflect.Array, reflect.String, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr, reflect.Float32, reflect.Float64, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
			return true
		}
	}
	return false
}

func (m *Max) validate(v *Validator, maxI interface{}, msg ...string) {
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
}
