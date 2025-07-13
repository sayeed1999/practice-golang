package namedresultparams

import "fmt"

// Note: - When Named result paramters are better?
// Clarify purpose of return values which can be useful in complex functions.
// - They can make the code more readable by providing context to the return values.

func GetStatus() (code int, message string) {
	code = 200
	message = "OK"
	return // returns code, message implicitly
}

func Divide(a, b int) (result int, err error) {
	if b == 0 {
		err = fmt.Errorf("division by zero")
	} else {
		result = a / b
	}
	return // returns result, err implicitly
}
