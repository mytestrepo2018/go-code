package main

import (
	"fmt"
	"strings"
)

const englishHelloPrefix = "hello "
const spanishHelloPrefix = "hola "
const frenchHelloPrefix = "bonjour "

func greetingPrefix(language string) string {
	p := strings.ToLower(language)
	switch p {
	case "spanish":
		p = spanishHelloPrefix
	case "french":
		p = frenchHelloPrefix
	case "english":
	default:
		p = englishHelloPrefix
	}
	return string(p)
}

// Hello takes a name and language and prints a welcome message
func Hello(name string, language string) string {
	if len(name) == 0 {
		name = "world"
	}

	return greetingPrefix(language) + name + "!"
}

func main() {
	fmt.Println(Hello("", "french"))
}
