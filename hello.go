package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPefix

	switch language {
	case spanish:
		return spanishHelloPrefix + name
	case french:
		return frenchHelloPrefix + name
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
