package main

import "fmt"

const englishHelloPrefix = "Hello, "

// Hello returns the string >> Hello, world <<
func Hello() string {
	return "Hello, world"
}

// HelloYou returns the string >> Hello, world <<
func HelloYou(name string) string {
	if name == "" {
		name = "world"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(HelloYou("Maria"))
}
