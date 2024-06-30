package helloworld

import "fmt"

const (
	spanish            = "Spanish"
	french             = "French"
	englishHelloPefix  = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

// Hello prints greeting in different languages
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name

}

// greetingPrefix returns the greeting prefix by language
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
