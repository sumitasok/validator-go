package validator

import (
	"errors"
	// "fmt"
	"reflect"
)

type Validator struct {
	Object interface{}
	Errors []error
}

func (v *Validator) Required() *Validator {
	if v.Object == nil {
		v.Add("required")
	}

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
	}

	// obj := dereference(objPtr)
	// if timeVal, ok := obj.(time.Time); ok {
	// 	return !timeVal.IsZero()
	// }

	// val := reflect.ValueOf(obj)
	// return !IsZero(val)

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
