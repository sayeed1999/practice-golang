package main

import "strings"

func ConvertToRomain(num int) string {
	var result strings.Builder

	for num > 0 {
		if num >= 5 {
			result.WriteString("V")
			num -= 5
		} else if num >= 4 {
			result.WriteString("IV")
			num -= 4
		} else {
			result.WriteString("I")
			num--
		}
	}

	return result.String()
}
