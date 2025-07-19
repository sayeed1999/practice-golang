package main

import "fmt"

const (
	spanish            = "Spanish"
	french             = "French"
	englishHelloPrefix = "Hello"
	spanishHelloPrefix = "Hola"
	frenchHelloPrefix  = "Bonjour"
)

func main() {
	fmt.Println(Hello())
}

func Hello() string {
	return englishHelloPrefix + ", World!"
}

func HelloName(name string, language string) string {
	if name == "" {
		name = "Sayeed"
	}

	return fmt.Sprintf("%s, %s!", greetingPrefix(language), name)
}

// By named return value, we make it more clear what the function returns.
func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
