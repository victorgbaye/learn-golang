package main

import "fmt"
const englishHelloPrefix string = "Hello, "
const spanishHelloPrefix string = "Hola, "
const frenchHelloPrefix string = "Bonjour, "

func Hello(name, language string) string{
	if name == ""{
		return englishHelloPrefix + "world"
	}
	
	if language == "Spanish" {
		return spanishHelloPrefix + name
	}
	if language == "French" {
		return frenchHelloPrefix + name
	}
	if language == "English" {
		return englishHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main () {
	fmt.Println(Hello("Chris", "English"))
}