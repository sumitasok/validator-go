package validator

import (
	"reflect"
	"strconv"
)

var (
	tArray  = []string{"Lorem", "ipsum", "dolor", "sit", "amet", "consectetur", "adipiscing"} // reflect.Slice
	tString = "Lorem ipsum dolor sit amet, consectetur adipiscing"                            // reflect.String
	tInt    = int(127)                                                                        // reflect.Int

	tprobabs = tProbabilities{
		tObjType{tArray,
			tOps{"maxFail": tProp{V: int(4), E: "maximum 4 numbers allowed", A: true}},
		},
		tObjType{tString,
			tOps{"maxFail": tProp{V: int(7), E: "maximum 7 characters allowed", A: true}},
		},
		tObjType{tInt,
			tOps{"maxFail": tProp{V: int8(7), E: "maximum 7 is allowed", A: true}},
		},
	}
)

type tObjType struct {
	Obj interface{}
	Ops tOps
}

func (o tObjType) String(i int, op string, using interface{}) string {
	return "in condition: (" + strconv.Itoa(i+1) + ") op(" +
		op + ") type(" + reflect.ValueOf(o.Obj).Kind().String() +
		")" + " compared using (" + reflect.ValueOf(using).Kind().String()
}

type tOps map[string]tProp

type tProp struct {
	V interface{} // Value
	E string      // Error message
	A bool        // Allowed
}

type tProbabilities []tObjType
