package main

import "strings"

// Inefficient implementation of Repeat function
// Strings in go are immutable, so concatenating strings in a loop is inefficient.

func Repeat(character string, count int) string {
	repeated := ""
	for i := 0; i < count; i++ {
		repeated += character
	}
	return repeated
}

func RepeatFast(character string, count int) string {
	var repeated strings.Builder
	for i := 0; i < count; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
