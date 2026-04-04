package reflection

import "reflect"

// golang challenge: write a function `walk(x interface{}, fn func(string))`
// which takes a struct `x` and calls `fn` for all string fields found inside.
// difficulty level: recursively.

// provided requirement: interface{} is a struct
func Walk(x interface{}, fn func(string)) {
	val := reflect.ValueOf(x)

	// first extract the value if val is pointing to a Ptr
	if val.Kind() == reflect.Pointer {
		val = val.Elem() // necessary
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		// checking if the field Kind is a string
		if field.Kind() == reflect.String {
			fn(field.String())
		}

		// do recursion if the field itself is a struct !!
		if field.Kind() == reflect.Struct {
			Walk(field.Interface(), fn) // VVI: must parse the field as an interface{} here!
		}
	}
}
