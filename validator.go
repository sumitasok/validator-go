package validator

import (
	// "fmt"
	"errors"
	"reflect"
)

type Typer interface {
	Required() bool
}

type Validator struct {
	Object Typer
	Error  error
}

func (v *Validator) Required() *Validator {
	if v.Object.Required() == false {
		v.Error = errors.New("required field")
	}
	return v
}

func (v *Validator) On(obj interface{}) *Validator {

	objType := reflect.TypeOf(obj).String()
	switch objType {
	case "string":
		if str, ok := obj.(string); ok {
			v.Object = VString{Data: str}
		}
	}
	return v
}

type VString struct {
	Data string
}

func (v VString) Required() bool {
	if v.Data == "" {
		return false
	} else {
		return true
	}
}
