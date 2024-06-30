package main

import "fmt"

const englishHelloPefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
