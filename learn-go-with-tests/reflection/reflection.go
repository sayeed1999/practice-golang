package reflection

import "reflect"

// golang challenge: write a function `walk(x interface{}, fn func(string))`
// which takes a struct `x` and calls `fn` for all string fields found inside.
// difficulty level: recursively.

// provided requirement: interface{} is a struct
func Walk(x interface{}, fn func(string)) {
	val := reflect.ValueOf(x)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fn(field.String()) // fix
	}
}
