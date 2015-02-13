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

func typeAllow(a interface{}, b interface{}) bool {
	aKind := reflect.ValueOf(a).Kind()
	bKind := reflect.ValueOf(b).Kind()
	aType := reflect.ValueOf(a).Type()
	bType := reflect.ValueOf(b).Type()
	if aKind == reflect.Struct && bKind != reflect.Struct {
		return false
	}

	if aKind != reflect.Struct && bKind == reflect.Struct {
		return false
	}

	//	If both are struct
	if aKind == reflect.Struct && bKind == reflect.Struct {
		if aType == bType {
			if aType == reflect.ValueOf(time.Time{}).Type() {
				// only time is implemented yet
				return true
			}
			return false
		}
		return false
	}

	if aKind == reflect.String {
		if bKind == reflect.Float64 {
			return false
		}
	}

	return true
}

func (m *Max) validate(v *Validator, maxI interface{}, msg ...string) {
	if !typeAllow(v.Object, maxI) {
		v.Add("cannot be applied on this object")
		return
	}

	fmt.Println(reflect.ValueOf(v.Object).Kind())

	switch reflect.ValueOf(v.Object).Kind() {
	case reflect.Array, reflect.String:
		max := reflect.ValueOf(maxI).Int()
		if int64(reflect.ValueOf(v.Object).Len()) > max {
			defaultMsg := fmt.Sprintf("maximum %s characters allowed", strconv.Itoa(int(max)))
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
