package validator

import (
	"fmt"
	"reflect"
	"time"
)

type Max struct {
	V *Validator
}

func (m *Max) Allow(obj interface{}) bool {
	objType := reflect.ValueOf(obj).Kind()
	valueType := reflect.ValueOf(m.V.Object).Kind()
	fmt.Println(objType, valueType)
	switch objType {
	case reflect.Int:
		switch valueType {
		case reflect.Array, reflect.String, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr, reflect.Float32, reflect.Float64, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
			return true
		}
	}
	return false
}

func typeAllow(a interface{}, b interface{}) bool {
	if reflect.ValueOf(a).Kind() == reflect.Struct && reflect.ValueOf(b).Kind() != reflect.Struct {
		return false
	}

	if reflect.ValueOf(a).Kind() != reflect.Struct && reflect.ValueOf(b).Kind() == reflect.Struct {
		return false
	}

	return true
}

func (m *Max) validate(v *Validator, maxI interface{}, msg ...string) {
	if !typeAllow(maxI, v.Object) {
		v.Add("cannot be applied on this object")
		return
	}

	switch reflect.ValueOf(v.Object).Kind() {
	case reflect.Array, reflect.String:
		max := reflect.ValueOf(maxI).Int()
		if int64(reflect.ValueOf(v.Object).Len()) > max {
			defaultMsg := fmt.Sprintf("maximum %x characters allowed", max)
			v.Add(defaultMsg, msg...)
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		max := reflect.ValueOf(maxI).Int()
		if reflect.ValueOf(v.Object).Int() > max {
			defaultMsg := fmt.Sprintf("maximum %x is allowed", max)
			v.Add(defaultMsg, msg...)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		max := reflect.ValueOf(maxI).Uint()
		if reflect.ValueOf(v.Object).Uint() > max {
			defaultMsg := fmt.Sprintf("maximum %x is allowed", max)
			v.Add(defaultMsg, msg...)
		}
	case reflect.Float32, reflect.Float64:
		max := reflect.ValueOf(maxI).Float()
		if reflect.ValueOf(v.Object).Float() > max {
			defaultMsg := fmt.Sprintf("maximum %x is allowed", max)
			v.Add(defaultMsg, msg...)
		}
	case reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		max := reflect.ValueOf(maxI).Int()
		if int64(reflect.ValueOf(v.Object).Len()) > max {
			defaultMsg := fmt.Sprintf("maximum %x numbers allowed", max)
			v.Add(defaultMsg, msg...)
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
