package main

import (
	"fmt"
)
const englishHelloPrefix string = "Hello, "
const spanishHelloPrefix string = "Hola, "
const frenchHelloPrefix string = "Bonjour, "

func Hello(name, language string) string{
	if name == ""{
		return englishHelloPrefix + "world"
	}

	 prefix := englishHelloPrefix

	switch language{
	case "Spanish":
		prefix = spanishHelloPrefix
	case "French":
		prefix = frenchHelloPrefix
	}
	return prefix + name
}

func main () {
	fmt.Println(Hello("Chris", "English"))
}