package main

import "fmt"

const englishHelloPefix = "Hello, "

func Hello(name string) string {
	return englishHelloPefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
