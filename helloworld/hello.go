package main

import "fmt"

const (
	englishHelloPrefix = "Hello,"
	spanishHelloPrefix = "Hola,"
	frenchHelloPrefix  = "Bonjour,"

	spanishLang = "Spanish"
	frenchLang  = "French"
)

func main() {
	fmt.Println(Hello("Name", ""))
}

func Hello(name, language string) string {
	prefix := getPrefix(language)

	if len(name) == 0 {
		return fmt.Sprintf("%s World!", prefix)
	}
	return fmt.Sprintf("%s %s!", prefix, name)
}

func getPrefix(language string) string {
	switch language {
	case spanishLang:
		return spanishHelloPrefix
	case frenchLang:
		return frenchHelloPrefix
	default:
		return englishHelloPrefix
	}
}
