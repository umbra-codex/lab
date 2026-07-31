package main

import "fmt"

const (
	englishHelloPrefix = "Hello"
	frenchHelloPrefix  = "Bonjour"
	spanishHelloPrefix = "Hola"
	french             = "French"
	spanish            = "Spanish"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == "French" {
		return fmt.Sprintf("%s %s!", frenchHelloPrefix, name)
	}

	if language == "Spanish" {
		return fmt.Sprintf("%s %s!", spanishHelloPrefix, name)
	}

	return fmt.Sprintf("%s %s!", englishHelloPrefix, name)
}

func main() {
	fmt.Println(Hello("", ""))
}
