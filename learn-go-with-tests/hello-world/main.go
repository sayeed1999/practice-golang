package main

import "fmt"

func Hello() string {
	return "Hello, World!"
}

func HelloName(name string) string {
	if name == "" {
		return "Hello, Sayeed!"
	}

	return "Hello, " + name + "!"
}

func main() {
	fmt.Println(Hello())
}
