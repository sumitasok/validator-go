package validator

import (
	// "fmt"
	"reflect"
)

type Validator struct {
	Object Typer
}

type Typer interface {
}

func (v *Validator) Work(obj interface{}) {

	objType := reflect.TypeOf(obj).String()
	switch objType {
	case "string":
		if str, ok := obj.(string); ok {
			v.Object = VString{str}
		}
	}

}

type VString struct {
	Data string
}
