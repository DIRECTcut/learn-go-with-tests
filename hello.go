package main

import "fmt"

const englishHelloPrefix = "Hello,"

func main() {
	fmt.Println(Hello("Name"))
}

func Hello(name string) string {
	return fmt.Sprintf("%s %s!", englishHelloPrefix, name)
}
