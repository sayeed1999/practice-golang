package main

import "strings"

func ConvertToRomain(num int) string {
	var result strings.Builder

	for num > 0 {
		if num >= 50 {
			result.WriteString("L")
			num -= 50
		} else if num >= 40 {
			result.WriteString("XL")
			num -= 40
		} else if num >= 10 {
			result.WriteString("X")
			num -= 10
		} else if num >= 9 {
			result.WriteString("IX")
			num -= 9
		} else if num >= 5 {
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
