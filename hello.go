package main

import "fmt"

const (
	spanish            = "Spanish"
	french             = "French"
	englishHelloPefix  = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name

}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPefix
	}

	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
