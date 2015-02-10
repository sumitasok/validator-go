package validator

import (
	"fmt"
	"reflect"
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
