package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

const spanish = "Spanish"
const french = "French"

// Hello returns the string >> Hello, world <<
func Hello() string {
	return "Hello, world"
}

// HelloYou returns the string >> Hello, world <<
func HelloYou(name, language string) string {
	if name == "" {
		name = "World"
	}
	var prefix string

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(HelloYou("Maria", "Portuguese"))
}
