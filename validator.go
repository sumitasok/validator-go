package validator

import (
	// "fmt"
	"errors"
	// "reflect"
)

type Validator struct {
	Object interface{}
	Errors []error
}

func (v *Validator) Required() *Validator {
	if v.Object == nil {
		v.Add("required")
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
