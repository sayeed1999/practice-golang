package main

import "strings"

func ConvertToRoman(num int) string {
	var result strings.Builder

	for num > 0 {
		if num >= 1000 {
			result.WriteString("M")
			num -= 1000
		} else if num >= 900 {
			result.WriteString("CM")
			num -= 900
		} else if num >= 500 {
			result.WriteString("D")
			num -= 500
		} else if num >= 400 {
			result.WriteString("CD")
			num -= 400
		} else if num >= 100 {
			result.WriteString("C")
			num -= 100
		} else if num >= 90 {
			result.WriteString("XC")
			num -= 90
		} else if num >= 50 {
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
