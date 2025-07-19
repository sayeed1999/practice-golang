package main

import "fmt"

const (
	spanish            = "Spanish"
	french             = "French"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello() string {
	return englishHelloPrefix + "World!"
}

func HelloName(name string, language string) string {
	if name == "" {
		name = "Sayeed"
	}

	if language == spanish {
		return spanishHelloPrefix + name + "!"
	}

	if language == french {
		return frenchHelloPrefix + name + "!"
	}

	return englishHelloPrefix + name + "!"
}

func main() {
	fmt.Println(Hello())
}
