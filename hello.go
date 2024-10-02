package main

import "fmt"

const englishHelloPrefix = "Hello,"

func main() {
	fmt.Println(Hello("Name"))
}

func Hello(name string) string {
	if len(name) == 0 {
		return fmt.Sprintf("%s World!", englishHelloPrefix)
	}
	return fmt.Sprintf("%s %s!", englishHelloPrefix, name)
}
