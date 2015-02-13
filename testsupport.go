package validator

import (
	"time"
)

var (
	tArray  = []string{"Lorem", "ipsum", "dolor", "sit", "amet", "consectetur", "adipiscing"}                                                                                                                                                                                  // reflect.Slice
	tString = "Lorem ipsum dolor sit amet, consectetur adipiscing Lorem ipsum dolor sit amet, consectetur adipiscing Lorem ipsum dolor sit amet, consectetur adipiscing Lorem ipsum dolor sit amet, consectetur adipiscing Lorem ipsum dolor sit amet, consectetur adipiscing" // reflect.String

	tInt    = int(127)   // reflect.Int
	tFloat  = 1.1        // reflect.Float
	tTime   = time.Now() // time.Time
	tStruct = testStr{}

	tprobabs = tProbabilities{
		tObjType{tArray,
			tOps{"maxFail": []tProp{
				tProp{V: int(4), E: "maximum 4 numbers allowed", A: true},
			}},
		},
		tObjType{tString,
			tOps{"maxFail": []tProp{
				tProp{V: int(7), E: "maximum 7 characters allowed", A: true},
				tProp{V: int16(128), E: "maximum 128 characters allowed", A: true},
				tProp{V: int32(128), E: "maximum 128 characters allowed", A: true},
				tProp{V: int64(128), E: "maximum 128 characters allowed", A: true},
			}},
		},
		tObjType{tInt,
			tOps{
				"maxFail": []tProp{
					tProp{V: int8(7), E: "maximum 7 is allowed", A: true},
				},
				"allow": []tProp{
					tProp{V: tTime, A: false},
					tProp{V: int(7), A: true},
					tProp{V: tStruct, A: false},
					tProp{V: int64(7), A: true},
				},
			},
		},
		tObjType{tFloat,
			tOps{
				"maxFail": []tProp{
					tProp{V: tTime, E: "cannot be applied on this object", A: false},
				},
				"allow": []tProp{
					tProp{V: tTime, A: false},
					tProp{V: int(7), A: true},
					tProp{V: tStruct, A: false},
				},
			},
		},
		tObjType{tTime,
			tOps{
				"allow": []tProp{
					tProp{V: tTime, A: true},
					tProp{V: int(7), A: false},
					tProp{V: tStruct, A: false},
					tProp{V: int64(7), A: false},
				},
			},
		},
	}
)

type tObjType struct {
	Obj interface{}
	Ops tOps
}

type tOps map[string]([]tProp)

type tProp struct {
	V interface{} // Value
	E string      // Error message
	A bool        // Allowed
}

type tProbabilities []tObjType
