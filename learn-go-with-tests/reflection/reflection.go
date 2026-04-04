package reflection

import "reflect"

// golang challenge: write a function `walk(x interface{}, fn func(string))`
// which takes a struct `x` and calls `fn` for all string fields found inside.
// difficulty level: recursively.

// provided requirement: interface{} is a struct
func Walk(x interface{}, fn func(string)) {
	val := getValue(x)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		switch field.Kind() {
		case reflect.String:
			fn(field.String()) // call the function if it is a string
		case reflect.Struct: // do recursion if the field itself is a struct !!
			Walk(field.Interface(), fn) // VVI: must parse the field as an interface{} here!

		}
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
