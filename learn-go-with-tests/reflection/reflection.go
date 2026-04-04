package reflection

import "reflect"

// golang challenge: write a function `walk(x interface{}, fn func(string))`
// which takes a struct `x` and calls `fn` for all string fields found inside.
// difficulty level: recursively.

// provided requirement: interface{} is a struct
func Walk(x interface{}, fn func(string)) {
	val := getValue(x)

	switch val.Kind() {
	// if [], need to walk over indices and recurse
	case reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			Walk(val.Index(i).Interface(), fn) // VVI: must parse the field as an interface
		}
	// if {}, need to walk over fields and recurse
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			Walk(val.Field(i).Interface(), fn) // VVI: must parse the field as an interface
		}
	// if string, call the desired function!
	case reflect.String:
		fn(val.String())
	default:
		// else, do nothing
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	// first extract the value if val is pointing to a Ptr
	if val.Kind() == reflect.Pointer {
		val = val.Elem() // necessary
	}

	return val
}
